package libtc

import (
	"errors"
	"sync"
)

type TrafficFilter struct {
	SrcAddr string `yaml:"src_addr" json:"src_addr" bson:"src_addr"`
	SrcMask string `yaml:"src_mask" json:"src_mask" bson:"src_mask"`
	SrcPort int    `yaml:"src_port" json:"src_port" bson:"src_port"`
	DstAddr string `yaml:"dst_addr" json:"dst_addr" bson:"dst_addr"`
	DstMask string `yaml:"dst_mask" json:"dst_mask" bson:"dst_mask"`
	DstPort int    `yaml:"dst_port" json:"dst_port" bson:"dst_port"`
	ToS     int    `yaml:"tos" json:"tos" bson:"tos"`
}

type RateType string

const (
	KbitSec  RateType = "Kbit"
	MbitSec           = "Mbit"
	GbitSec           = "Gbit"
	KbyteSec          = "kbps"
	MbyteSec          = "mbps"
	GbyteSec          = "gbps"
)

type QosItem struct {
	Filters  []TrafficFilter `yaml:"filters" json:"filters" bson:"filters"`
	Block    bool            `yaml:"block" json:"block" bson:"block"`
	Rate     int             `yaml:"rate" json:"rate" bson:"rate"`
	RateType RateType        `yaml:"rate_type" json:"rate_type" bson:"rate_type"` // what-per-second
}

type UserClass struct {
	Name   string    `yaml:"name" json:"name" bson:"name"`
	QosIn  []QosItem `yaml:"qos_in" json:"qos_in" bson:"qos_in"`
	QosOut []QosItem `yaml:"qos_out" json:"qos_out" bson:"qos_out"`
}

type User struct {
	Address string
	Class   string
	Ifaces  []string
}

type innerRes struct {
	filters map[int]bool
	classes map[int]bool
}

type ifaceState struct {
	users map[string]innerRes
	qdisc map[int]innerRes
	dir   int
}

const (
	DIR_IN   = 1
	DIR_OUT  = -1
	DIR_NONE = 0
)

type TcBind interface {
	AddFilter(dev *Iface, f Filter)
	AddClass(dev *Iface, c Class)
	AddQdisc(dev *Iface, qdisc QDisc)
	DelFilter(dev *Iface, f Filter)
	DelClass(dev *Iface, c Class)
	DelQdisc(dev *Iface, qdisc QDisc)
	Get()

	Prepare()
	Commit() error
	Error(err error)

	Init(db *Database)
}

type Database struct {
	users   map[string]User
	ifaces  map[string]Iface
	classes map[string]UserClass
	states  map[string]ifaceState
	mutex   sync.Mutex
	ready   bool
	Tc      TcBind
}

func (db *Database) SetInterfaces(ifaces map[string]Iface) {
	db.ifaces = ifaces
	db.parseInterfaces()
}

func (db *Database) Load() error {
	return db.init()
}

func (db *Database) Close() error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Tc.Prepare()

	for uname, user := range db.users {
		user.Ifaces = []string{}
		db.users[uname] = user
	}

	for ifname, _ := range db.ifaces {
		db.deinitIface(ifname)
	}
	return db.commit()
}

func (db *Database) InitIface(iface string, dir int, clsname string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if !db.ready {
		return errors.New("Data dont loaded")
	}
	if _, ok := db.ifaces[iface]; !ok {
		return errors.New("Interface doesnt exist")
	}
	if _, ok := db.classes[clsname]; !ok {
		return errors.New("Default class doesnt exist")
	}
	db.Tc.Prepare()
	db.initIface(iface, dir, clsname)
	return db.commit()
}

func (db *Database) DeinitIface(ifname string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if !db.ready {
		return errors.New("Data dont loaded")
	}
	if _, ok := db.ifaces[ifname]; !ok {
		return errors.New("Interface doesnt exist")
	}
	db.Tc.Prepare()
	for uname, user := range db.users {
		var newifaces []string
		for _, iface := range user.Ifaces {
			if iface != ifname {
				newifaces = append(newifaces, iface)
			}
		}
		user.Ifaces = newifaces
		db.users[uname] = user
	}

	db.deinitIface(ifname)
	return db.commit()
}

func (db *Database) NewUser(mac string, class string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if !db.ready {
		return errors.New("Data dont loaded")
	}
	mac = normalizeMac(mac)
	var u, ok = db.users[mac]
	if ok {
		if u.Class != class {
			return errors.New("User exists with another class")
		}
		return nil
	}
	if _, ok := db.classes[class]; !ok {
		return errors.New("Class not found")
	}
	db.users[mac] = User{
		Address: mac,
		Class:   class,
	}
	return nil
}

func (db *Database) AssignUser(iface_in, iface_out string, mac string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if !db.ready {
		return errors.New("Data dont loaded")
	}
	mac = normalizeMac(mac)
	var u, ok = db.users[mac]
	if !ok {
		return errors.New("User doesnt exist")
	}
	var in, out bool
	if ok {
		for _, i := range u.Ifaces {
			if i == iface_in {
				in = true
			}
			if i == iface_out {
				out = true
			}
		}
		if in && out {
			return nil
		}
		if !in {
			u.Ifaces = append(u.Ifaces, iface_in)
		}
		if !out {
			u.Ifaces = append(u.Ifaces, iface_out)
		}
	} else {
		return errors.New("User doesnt exist")
	}
	db.Tc.Prepare()
	db.addUser(iface_in, mac, u.Class, DIR_IN)
	db.addUser(iface_out, mac, u.Class, DIR_OUT)
	db.users[mac] = u
	return db.commit()
}

func (db *Database) DeassignUser(iface string, mac string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if !db.ready {
		return errors.New("Data dont loaded")
	}
	mac = normalizeMac(mac)
	var u, ok = db.users[mac]
	if !ok {
		return errors.New("User not found")
	}
	db.Tc.Prepare()
	for index, i := range u.Ifaces {
		if i == iface {
			u.Ifaces = append(u.Ifaces[:index], u.Ifaces[index+1:]...)
			if len(u.Ifaces) == 0 {
				delete(db.users, mac)
			} else {
				db.users[mac] = u
			}
			db.delUser(iface, mac)
		}
	}
	return db.commit()
}

func (db *Database) DelUser(mac string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if !db.ready {
		return errors.New("Data dont loaded")
	}
	mac = normalizeMac(mac)
	var u, ok = db.users[mac]
	if !ok {
		return errors.New("User not found")
	}
	db.Tc.Prepare()
	for _, i := range u.Ifaces {
		db.delUser(i, mac)
	}
	delete(db.users, mac)
	return db.commit()
}

func (db *Database) ChangeUser(mac string, class string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if !db.ready {
		return errors.New("Data dont loaded")
	}
	mac = normalizeMac(mac)
	var u, ok = db.users[mac]
	if !ok {
		return errors.New("User doesnt exist")
	}
	if u.Class == class {
		return nil
	}
	db.Tc.Prepare()
	for _, iface := range u.Ifaces {
		db.changeUser(iface, mac, class)
	}
	return db.commit()
}

func (db *Database) AddClass(class UserClass) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if !db.ready {
		return errors.New("Data dont loaded")
	}
	if _, ok := db.classes[class.Name]; ok {
		return errors.New("Class already exists")
	}
	if db.classes == nil {
		db.classes = map[string]UserClass{}
	}
	db.classes[class.Name] = class
	return nil
}

func (db *Database) DelClass(classname string, withusers bool) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if !db.ready {
		return errors.New("Data dont loaded")
	}
	var _, ok = db.classes[classname]
	if !ok {
		return nil
	}
	var users []string
	for mac, user := range db.users {
		if user.Class == classname {
			users = append(users, mac)
		}
	}
	if !withusers && len(users) != 0 {
		return errors.New("Dont del class with users")
	}
	db.Tc.Prepare()
	for _, mac := range users {
		mac = normalizeMac(mac)
		var user = db.users[mac]
		for _, i := range user.Ifaces {
			db.delUser(i, mac)
		}
		delete(db.users, mac)
	}

	delete(db.classes, classname)
	return db.commit()
}

func (db *Database) ChangeClass(class UserClass, withusers bool) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if !db.ready {
		return errors.New("Data dont loaded")
	}
	// TODO
	return errors.New("Not implemented")
	db.Tc.Prepare()
	return db.commit()
}
