package libtc

import (
	"errors"
	"sync"
)

type TrafficFilter struct {
	SrcAddr string `yaml:"src_addr"`
	SrcMask string `yaml:"src_mask"`
	SrcPort int    `yaml:"src_port"`
	DstAddr string `yaml:"dst_addr"`
	DstMask string `yaml:"dst_mask"`
	DstPort int    `yaml:"dst_port"`
	ToS     int    `yaml:"tos"`
}

type TrafficClass struct {
	Name    string          `yaml:"name"`
	Filters []TrafficFilter `yaml:"filters"`
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
	Class    TrafficClass `yaml:"class"`
	Block    bool         `yaml:"block"`
	Rate     int          `yaml:"rate"`
	RateType RateType     `yaml:"rate_type"` // what-per-second
}

type UserClass struct {
	Name   string    `yaml:"name"`
	QosIn  []QosItem `yaml:"qos_in"`
	QosOut []QosItem `yaml:"qos_out"`
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
	users map[int]innerRes
	qdisc map[int]innerRes
	dir   int
}

const (
	DIR_IN   = 1
	DIR_OUT  = -1
	DIR_NONE = 0
)

type Database struct {
	users   map[string]User
	ifaces  map[string]Iface
	classes map[string]UserClass
	states  map[string]ifaceState
	mutex   sync.Mutex
	ready   bool
}

func (db *Database) Load() error {
	return db.init()
}

func (db *Database) Close() error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	for uname, user := range db.users {
		user.Ifaces = []string{}
		db.users[uname] = user
	}

	for ifname, _ := range db.ifaces {
		if err := db.deinitIface(ifname); err != nil {
			return err
		}
	}
	return nil
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
	return db.initIface(iface, dir, clsname)
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

	return db.deinitIface(ifname)
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
	if err := db.addUser(iface_in, mac, u.Class, DIR_IN); err != nil {
		return err
	}
	if err := db.addUser(iface_out, mac, u.Class, DIR_OUT); err != nil {
		return err
	}
	db.users[mac] = u
	return nil
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
	for index, i := range u.Ifaces {
		if i == iface {
			u.Ifaces = append(u.Ifaces[:index], u.Ifaces[index+1:]...)
			if len(u.Ifaces) == 0 {
				delete(db.users, mac)
			} else {
				db.users[mac] = u
			}
			return db.delUser(iface, mac)
		}
	}
	return nil
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
	for _, i := range u.Ifaces {
		if err := db.delUser(i, mac); err != nil {
			return err
		}
	}
	delete(db.users, mac)
	return nil
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
	for _, iface := range u.Ifaces {
		if err := db.changeUser(iface, mac, class); err != nil {
			return err
		}
	}
	return nil
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
	for _, mac := range users {
		mac = normalizeMac(mac)
		var user = db.users[mac]
		for _, i := range user.Ifaces {
			if err := db.delUser(i, mac); err != nil {
				return err
			}
		}
		delete(db.users, mac)
	}

	delete(db.classes, classname)
	return nil
}

func (db *Database) ChangeClass(class UserClass, withusers bool) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	if !db.ready {
		return errors.New("Data dont loaded")
	}
	// TODO
	return errors.New("Not implemented")
}
