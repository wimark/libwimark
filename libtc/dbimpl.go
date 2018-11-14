package libtc

import (
	"errors"
	"fmt"
)

type innerDb struct {
}

const (
	DEFAULT_TABLE_HANDLE  = 0x800
	INGRESS_DISC_HANDLE   = 0xffff
	ROOT_DISC_HANDLE      = 0x8001
	ROOT_TABLE_HANDLE     = 0x7ff
	UNAUTH_CLASS_HANDLE   = 0xffff
	MAX_FILTER_NUM        = 0xfff
	UNAUTH_NETFILTER_MARK = 0x777
)

// TODO to settings
var (
	UNAUTH_DEFAULT_RATE = "1kbps"
	UNAUTH_DEFAULT_CEIL = "100gbps"
)

var MATCH_ALL_FILTER = FilterU32Match{}

var FAKE_QOS_OPTIONS = []Option{
	Option{Name: HTB_RATE_PARAM, Value: "1bps"},
	Option{Name: HTB_CEIL_PARAM, Value: "1bps"},
}
var UNLIM_QOS_OPTIONS = []Option{
	Option{Name: HTB_RATE_PARAM, Value: "100gbps"},
	Option{Name: HTB_CEIL_PARAM, Value: "100gbps"},
}

func (db *Database) init() error {

	var ifaces, err = GetIfaces()
	if err != nil {
		return err
	}

	*db = Database{
		classes: map[string]UserClass{},
		ifaces:  ifaces,
		states:  map[string]ifaceState{},
		users:   map[string]User{},
		ready:   true,
	}

	return nil
}

func (db *Database) initIface(ifname string, dir int, clsname string) error {

	if err := db.deinitIface(ifname); err != nil {
		return err
	}

	var iface = db.ifaces[ifname]
	var class = db.classes[clsname]
	var err error
	db.states[ifname] = ifaceState{
		users: map[int]innerRes{},
		qdisc: map[int]innerRes{},
		dir:   dir,
	}
	var qos []QosItem
	switch dir {
	case DIR_IN:
		qos = class.QosIn
	case DIR_OUT:
		qos = class.QosOut
	}

	err = cfgQdisc(db, &iface, splitMac("", dir).root)
	if err != nil {
		return err
	}

	var newqos QosItem
	if len(qos) == 0 {
		newqos = QosItem{Block: true}
	} else {
		// use the first one only for now
		// TODO make several classes maybe
		newqos = qos[0]
	}
	var acts []Action
	if newqos.Block {
		acts = append(acts, Action{Type: "drop"})
	}
	// unauth filter
	err = AddFilter(&iface, Filter{
		Type:     "u32",
		Protocol: "all",
		Spec: &FilterU32{
			Match: []FilterU32Match{
				FilterU32Match{
					Name:  "mark",
					Param: fmt.Sprintf("%#08x 0xffffffff", UNAUTH_NETFILTER_MARK),
				},
			},
		},
		Parent:  ROOT_DISC_HANDLE,
		ToClass: UNAUTH_CLASS_HANDLE,
		Actions: acts,
		Prio:    1,
	})
	if err != nil {
		return err
	}
	// unauth class
	err = AddClass(&iface, Class{
		Handle:     UNAUTH_CLASS_HANDLE,
		ParentDisc: ROOT_DISC_HANDLE,
		Options:    qos2rates(newqos),
	})
	if err != nil {
		return err
	}

	db.ifaces[ifname] = iface
	return nil
}

func (db *Database) deinitIface(ifname string) error {

	var iface = db.ifaces[ifname]
	var err error
	delete(db.states, ifname)

	// drop all (but ingress)
	// empty QDisc has zero parent => it's root
	if len(iface.Discs) != 0 {
		err = DelQdisc(&iface, QDisc{})
		if err != nil {
			return err
		}
		for key, val := range iface.Discs {
			if val.Type != INGRESS_VALUE {
				delete(iface.Discs, key)
			}
		}
	}
	db.ifaces[ifname] = iface
	return nil
}

func cfgQdisc(db *Database, iface *Iface, cfg qdiscCfgForMac) error {

	var err error
	var state = db.states[iface.Name]
	if state.qdisc == nil {
		state.qdisc = map[int]innerRes{}
		state.users = map[int]innerRes{}
	}
	if cfg.listId != -1 {
		state.users[cfg.listId] = innerRes{
			filters: map[int]bool{},
			classes: map[int]bool{},
		}
	}
	if _, ok := iface.Discs[cfg.disc]; ok {
		db.states[iface.Name] = state
		return nil
	}
	state.qdisc[cfg.disc] = innerRes{
		filters: map[int]bool{},
		classes: map[int]bool{},
	}

	// qdisc
	// tc qdisc add
	//	dev DEV
	//	parent PD:PC
	//	handle QDISC:
	//	htb
	//		default ffff
	err = AddQdisc(iface, QDisc{
		Type:        QDISC_HTB,
		Handle:      cfg.disc,
		ParentDisc:  cfg.pdisc,
		ParentClass: cfg.pclass,
		Options: []Option{
			Option{
				Name:  DEFAULT_PARAM,
				Value: fmt.Sprintf("%04x", UNAUTH_CLASS_HANDLE),
			},
		},
	})
	if err != nil {
		return err
	}
	// base filter table
	// tc filter add
	//	dev DEV
	//	parent QDISC:
	//	pref 100 protocol ip
	//	handle 7ff:
	//	u32
	//		divisor 256
	err = AddFilter(iface, Filter{
		Type:     "u32",
		Protocol: "ip",
		Spec: &FilterU32{
			Table: ROOT_TABLE_HANDLE,
			Div:   256,
		},
		Parent: cfg.disc,
		Prio:   100,
	})
	if err != nil {
		return err
	}
	// root filter
	// in: DEV, QDISC, LINK, HASH
	// tc filter add
	//	dev DEV
	//	parent QDISC:
	//	pref 100 protocol ip
	//	u32
	//		link 7ff:
	//		match u8 0x00 0x00 at 0
	//		hash HASH
	err = AddFilter(iface, Filter{
		Type:     "u32",
		Protocol: "ip",
		Spec: &FilterU32{
			LinkTable: ROOT_TABLE_HANDLE,
			Hash:      cfg.rhash,
		},
		Parent: cfg.disc,
		Prio:   100,
	})
	if err != nil {
		return err
	}
	db.states[iface.Name] = state
	return nil
}

func cfgClass(db *Database, iface *Iface, cfg qdiscCfgForMac, qos []QosItem) error {

	var err error
	var qdisc, ok = iface.Discs[cfg.disc]
	if !ok {
		return errors.New("No qdisc")
	}
	var state = db.states[iface.Name]

	// check table
	ok = false
	// tc filter add
	//	dev DEV
	//	parent QDISC:
	//	pref 100 protocol ip
	//	handle TABLE:
	//	u32
	//		divisor 256
	var table = Filter{
		Type:     "u32",
		Protocol: "ip",
		Spec: &FilterU32{
			Table: cfg.table,
			Div:   256,
		},
		Parent: cfg.disc,
		Prio:   100,
	}
	for _, filt := range qdisc.Filters {
		if filt.Spec.Handle() == table.Spec.Handle() {
			ok = true
			break
		}
	}
	if !ok {
		// create table
		err = AddFilter(iface, table)
		if err != nil {
			return err
		}
		//  create filter to table
		// tc filter add
		//	dev DEV
		//	parent QDISC:
		//	pref 100 protocol ip
		//	u32
		//		ht ROOT:
		//		link TABLE
		//		hash LHASH
		//		sample RSAMPLE
		//		match ALL
		err = AddFilter(iface, Filter{
			Type:     "u32",
			Protocol: "ip",
			Spec: &FilterU32{
				Table:     ROOT_TABLE_HANDLE,
				LinkTable: cfg.table,
				Hash:      cfg.lhash,
				Sample:    cfg.rsample,
			},
			Parent: cfg.disc,
			Prio:   100,
		})
		if err != nil {
			return err
		}
	}

	var newqos = qos
	if len(newqos) == 0 {
		newqos = []QosItem{QosItem{}}
	}

	var clsNum = cfg.class
	var filtNum = 1
	for _, qos := range newqos {
		if cfg.listId != -1 {
			for ; clsNum < cfg.classLim; clsNum++ {
				if _, ok := state.qdisc[cfg.disc].classes[clsNum]; !ok {
					break
				}
			}
			if clsNum >= cfg.classLim {
				return errors.New("Class limit")
			}
			state.qdisc[cfg.disc].classes[clsNum] = true
			state.users[cfg.listId].classes[clsNum] = true
		}
		if _, ok := qdisc.Classes[clsNum]; ok {
			continue
		}
		// make class
		// tc class add
		//	dev DEV
		//	parent QDISC:
		//	classid QDISC:CLASS+classNum
		//	htb
		//		rate RATE ceil CEIL
		err = AddClass(iface, Class{
			Handle:     clsNum,
			ParentDisc: cfg.disc,
			Options:    qos2rates(qos),
		})
		if err != nil {
			return err
		}
		var newflt []TrafficFilter
		if len(qos.Class.Filters) == 0 {
			newflt = []TrafficFilter{TrafficFilter{}}
		} else {
			newflt = qos.Class.Filters
		}
		var acts []Action
		if qos.Block {
			acts = append(acts, Action{Type: "drop"})
		}
		for _, filter := range newflt {
			if cfg.listId != -1 {
				for ; filtNum <= MAX_FILTER_NUM; filtNum++ {
					if _, ok := state.qdisc[cfg.disc].filters[filtNum]; !ok {
						break
					}
				}
				if filtNum > MAX_FILTER_NUM {
					return errors.New("Filter limit")
				}
				state.qdisc[cfg.disc].filters[filtNum] = true
				state.users[cfg.listId].filters[filtNum] = true
			}
			// make filter
			// tc filter add
			//	dev DEV
			//	parent QDISC:
			//	pref 100 protocol ip
			//	handle ::num
			//	u32
			//		ht TABLE:
			//		match FILTER
			//		sample LSAMPLE
			//		classid QDISC:CLASS+
			err = AddFilter(iface, Filter{
				Type:     "u32",
				Protocol: "ip",
				Spec: &FilterU32{
					Table:     cfg.table,
					Index:     filtNum,
					LinkTable: -1,
					Sample:    cfg.lsample,
					Match:     append(filter2match(filter), cfg.match...),
				},
				Parent:  cfg.disc,
				ToClass: clsNum,
				Actions: acts,
				Prio:    100,
			})
			if err != nil {
				return err
			}
		}
	}
	db.states[iface.Name] = state

	return err
}

func qos2rates(qos QosItem) []Option {
	if qos.Rate == 0 {
		if qos.Block {
			return FAKE_QOS_OPTIONS
		} else {
			return UNLIM_QOS_OPTIONS
		}
	}
	return []Option{
		Option{Name: HTB_RATE_PARAM, Value: fmt.Sprintf("%d%s", qos.Rate, qos.RateType)},
		Option{Name: HTB_CEIL_PARAM, Value: fmt.Sprintf("%d%s", qos.Rate, qos.RateType)},
	}
}

func filter2match(flt TrafficFilter) []FilterU32Match {
	var res []FilterU32Match
	if len(flt.SrcAddr) != 0 {
		if len(flt.SrcMask) != 0 {
			res = append(res, FilterU32Match{
				Name:  "ip src",
				Param: fmt.Sprintf("%s/%s", flt.SrcAddr, flt.SrcMask),
			})
		} else {
			res = append(res, FilterU32Match{
				Name:  "ip src",
				Param: fmt.Sprintf("%s", flt.SrcAddr),
			})
		}
	}
	if len(flt.DstAddr) != 0 {
		if len(flt.DstMask) != 0 {
			res = append(res, FilterU32Match{
				Name:  "ip dst",
				Param: fmt.Sprintf("%s/%s", flt.DstAddr, flt.DstMask),
			})
		} else {
			res = append(res, FilterU32Match{
				Name:  "ip dst",
				Param: fmt.Sprintf("%s", flt.DstAddr),
			})
		}
	}
	if flt.SrcPort != 0 {
		res = append(res, FilterU32Match{
			Name:  "ip sport",
			Param: fmt.Sprintf("%d 0xffff", flt.SrcPort),
		})
	}
	if flt.DstPort != 0 {
		res = append(res, FilterU32Match{
			Name:  "ip dport",
			Param: fmt.Sprintf("%d 0xffff", flt.DstPort),
		})
	}
	if flt.ToS != 0 {
		res = append(res, FilterU32Match{
			Name:  "ip tos",
			Param: fmt.Sprintf("%s 0xff", hex(uint32(flt.ToS), 8)),
		})
	}
	return res
}

type qdiscCfgForMac struct {
	table    int
	disc     int
	class    int
	classLim int
	pdisc    int
	pclass   int
	rhash    FilterU32Match
	lhash    FilterU32Match
	rsample  FilterU32Match
	lsample  FilterU32Match
	match    []FilterU32Match
	listId   int
}

type macIndex struct {
	root qdiscCfgForMac
	leaf qdiscCfgForMac
}

func splitMac(mac string, dir int) macIndex {
	var octets = parseMac(mac)

	var index_part = octets[5] + (octets[4] << 8) + (octets[3] << 16)
	var list_part = octets[2] + (octets[1] << 8) + (octets[0] << 16)

	var root_table = index_part & 0xff
	var root_bucket = (index_part & 0xf00) >> 8
	var leaf_table = (index_part & 0xff0000) >> 16
	var leaf_bucket = (index_part & 0xf000) >> 12

	var rtFilt, rbFilt, ltFilt, lbFilt FilterU32Match
	var match []FilterU32Match

	switch dir {
	case DIR_IN:
		rtFilt = FilterU32Match{
			Value:  uint32(root_table),
			Mask:   mask(8, 0),
			Offset: -12,
		}
		rbFilt = FilterU32Match{
			Value:  uint32(root_bucket) << 8,
			Mask:   mask(4, 8),
			Offset: -12,
		}
		ltFilt = FilterU32Match{
			Value:  uint32(leaf_table) << 16,
			Mask:   mask(8, 16),
			Offset: -12,
		}
		lbFilt = FilterU32Match{
			Value:  uint32(leaf_bucket)<<12 + uint32(root_bucket)<<8,
			Mask:   mask(8, 8),
			Offset: -12,
		}
		match = []FilterU32Match{
			FilterU32Match{
				Value:  uint32(list_part&0xff) << 24,
				Mask:   mask(8, 24),
				Offset: -12,
			},
			FilterU32Match{
				Value:  uint32(list_part) >> 8,
				Mask:   mask(16, 0),
				Offset: -16,
			},
		}
	case DIR_OUT:
		rtFilt = FilterU32Match{
			Value:  uint32(root_table) << 16,
			Mask:   mask(8, 16),
			Offset: -4,
		}
		rbFilt = FilterU32Match{
			Value:  uint32(root_bucket) << 24,
			Mask:   mask(4, 24),
			Offset: -4,
		}
		ltFilt = FilterU32Match{
			Value:  uint32(leaf_table),
			Mask:   mask(8, 0),
			Offset: -8,
		}
		lbFilt = FilterU32Match{
			Value:  uint32(leaf_bucket)<<28 + uint32(root_bucket)<<24,
			Mask:   mask(8, 24),
			Offset: -4,
		}
		match = []FilterU32Match{FilterU32Match{
			Value:  uint32(list_part) << 8,
			Mask:   mask(24, 8),
			Offset: -8,
		}}
	}

	return macIndex{
		root: qdiscCfgForMac{
			table:   root_table,
			disc:    ROOT_DISC_HANDLE,
			class:   root_table + root_bucket<<8 + 1,
			pdisc:   0,
			pclass:  0,
			rhash:   rtFilt,
			lhash:   rbFilt,
			rsample: rtFilt,
			lsample: rbFilt,
			listId:  -1,
		},
		leaf: qdiscCfgForMac{
			disc:     root_table + root_bucket<<8 + 1,
			table:    leaf_table,
			class:    (leaf_table<<4+leaf_bucket)<<3 + 1,
			classLim: (leaf_table<<4+leaf_bucket+1)<<3 + 1,
			pdisc:    ROOT_DISC_HANDLE,
			pclass:   root_table + root_bucket<<8 + 1,
			rhash:    ltFilt,
			lhash:    lbFilt,
			rsample:  ltFilt,
			lsample:  lbFilt,
			match:    match,
			listId:   list_part,
		},
	}
}

func (db *Database) addUser(ifname, mac, clsname string, dir int) error {

	var err error
	var iface = db.ifaces[ifname]
	var qos []QosItem
	var index = splitMac(mac, dir)

	if i, ok := db.states[ifname]; ok && (i.dir != dir) {
		return errors.New("Trying to turn interface inside out")
	}
	switch dir {
	case DIR_IN:
		qos = db.classes[clsname].QosIn
	case DIR_OUT:
		qos = db.classes[clsname].QosOut
	}

	err = cfgClass(db, &iface, index.root, []QosItem{})
	if err != nil {
		return err
	}
	err = cfgQdisc(db, &iface, index.leaf)
	if err != nil {
		return err
	}
	err = cfgClass(db, &iface, index.leaf, qos)
	if err != nil {
		return err
	}
	db.ifaces[ifname] = iface

	return nil
}

func (db *Database) delUser(ifname, mac string) error {

	var err error
	var iface = db.ifaces[ifname]
	var state = db.states[ifname]
	var index = splitMac(mac, DIR_IN) // only hash/sample/match depend on dir
	var qdisc = iface.Discs[index.leaf.disc]
	var res = state.users[index.leaf.listId]

	var fltToRm []Filter
	for _, flt := range qdisc.Filters {
		if _, ok := res.classes[flt.ToClass]; ok {
			fltToRm = append(fltToRm, flt)
		}
	}

	for _, flt := range fltToRm {
		err = DelFilter(&iface, flt)
		if err != nil {
			return err
		}
	}
	for cls, _ := range res.classes {
		err = DelClass(&iface, qdisc.Classes[cls])
		if err != nil {
			return err
		}
	}

	db.ifaces[ifname] = iface
	db.states[ifname] = state

	return nil
}

func (db *Database) changeUser(ifname, mac, clsname string) error {

	if err := db.delUser(ifname, mac); err != nil {
		return err
	}
	if err := db.addUser(ifname, mac, clsname, db.states[ifname].dir); err != nil {
		return err
	}
	return nil
}
