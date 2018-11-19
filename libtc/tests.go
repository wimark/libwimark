package libtc

import (
	"errors"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"

	deep "github.com/go-test/deep"
)

const (
	test_dev_in  = "enp0s25"
	test_dev_out = "wlp3s0"

	user1 = "0123456789ab"
	user2 = "23456789abcd"
	user3 = "cdef456789ab"

	group1 = "patsak"
	group2 = "chatlan"
	group3 = "ecilop"
)

func ExecTest() error {

	var ifaces_old map[string]Iface
	var ifaces_new map[string]Iface
	var err error
	var tc TcBind = &ExecTcBind{}
	var db Database

	tc.Init(&db)

	fmt.Println("==== First get")

	tc.Prepare()
	tc.Get()
	if err := tc.Commit(); err != nil {
		return err
	}
	ifaces_old = db.ifaces
	var i, ok = ifaces_old[test_dev_in]
	if !ok {
		return errors.New("Test interface not found")
	}

	fmt.Println("==== Add stuff")

	tc.Prepare()
	// tc qdisc add
	//	dev enp0s25
	//	root
	//	handle 1:
	//	htb
	//		default 20
	tc.AddQdisc(&i, QDisc{
		Type:   "htb",
		Handle: 0x1,
		Options: []Option{
			Option{Name: "default", Value: "20"},
		},
	})
	// tc class add
	//	dev eth0
	//	parent 1:
	//	classid 1:1
	//	htb
	//		rate 90kbps ceil 90kbps
	tc.AddClass(&i, Class{
		Handle:      0x1,
		ParentDisc:  0x1,
		ParentClass: 0x0,
		Options: []Option{
			Option{Name: "rate", Value: "720Kbit"},
			Option{Name: "ceil", Value: "720Kbit"},
		},
	})
	// tc class add
	//	dev eth0
	//	parent 1:1
	//	classid 1:10
	//	htb
	//		rate 20kbps ceil 90kbps
	tc.AddClass(&i, Class{
		Handle:      0x10,
		ParentDisc:  0x1,
		ParentClass: 0x1,
		Options: []Option{
			Option{Name: "rate", Value: "160Kbit"},
			Option{Name: "ceil", Value: "560Kbit"},
		},
	})
	// tc class add
	//	dev eth0
	//	parent 1:1
	//	classid 1:20
	//	htb
	//		rate 70kbps ceil 90kbps
	tc.AddClass(&i, Class{
		Handle:      0x20,
		ParentDisc:  0x1,
		ParentClass: 0x1,
		Options: []Option{
			Option{Name: "rate", Value: "560Kbit"},
			Option{Name: "ceil", Value: "720Kbit"},
		},
	})
	// tc class add
	//	dev eth0
	//	parent 1:20
	//	classid 1:200
	//	htb
	//		rate 1kbps ceil 2kbps
	tc.AddClass(&i, Class{
		Handle:      0x200,
		ParentDisc:  0x1,
		ParentClass: 0x20,
		Options: []Option{
			Option{Name: "rate", Value: "8Kbit"},
			Option{Name: "ceil", Value: "16Kbit"},
		},
	})

	// tc qdisc add
	//	dev enp0s25
	//	parent 1:10
	//	handle 2:
	//	prio
	tc.AddQdisc(&i, QDisc{
		Type:        "prio",
		Handle:      0x2,
		ParentDisc:  0x1,
		ParentClass: 0x10,
		Options:     []Option{},
	})

	// tc qdisc add
	//	dev enp0s25
	//	ingress
	tc.AddQdisc(&i, QDisc{
		Type: INGRESS_VALUE,
	})

	// tc filter add
	//	dev enp0s25
	//	parent 1:
	//	prio 30
	//	protocol ip
	//	handle 1:
	//	u32
	//	divisor 256
	tc.AddFilter(&i, Filter{
		Type:     "u32",
		Protocol: "ip",
		Spec: &FilterU32{
			Table: 0x1,
			Div:   256,
		},
		Parent: 0x1,
		Prio:   30,
	})

	// tc filter add
	//	dev enp0s25
	//	parent 1:
	//	prio 10
	//	protocol 802_3
	//	handle ::1
	//	u32
	//		match ether src 12:34:56:78:9a:bc
	//	flowid 1:200
	tc.AddFilter(&i, Filter{
		Type:     "u32",
		Protocol: "802_3",
		Spec: &FilterU32{
			Index:     0x1,
			LinkTable: -1,
			Options: []Option{
				Option{Name: "match", Value: "ether src 12:34:56:78:9a:bc"},
			},
		},
		Parent:  0x1,
		ToClass: 0x200,
		Prio:    10,
		Actions: []Action{},
	})

	// tc filter add
	//	dev enp0s25
	//	parent 1:
	//	prio 20
	//	protocol ip
	//	handle ::2
	//	u32
	//		match ip dport 25 0xffff
	//	flowid 1:10
	tc.AddFilter(&i, Filter{
		Type:     "u32",
		Protocol: "ip",
		Spec: &FilterU32{
			Index:     0x2,
			LinkTable: -1,
			Options: []Option{
				Option{Name: "match", Value: "ip dport 25 0xffff"},
			},
		},
		Parent:  0x1,
		ToClass: 0x10,
		Prio:    20,
		Actions: []Action{},
	})

	// tc filter add
	//	dev enp0s25
	//	parent 2:
	//	prio 10
	//	protocol ip
	//	handle 1
	//	fw
	//	flowid 2:10
	tc.AddFilter(&i, Filter{
		Type:     "fw",
		Protocol: "ip",
		Spec: &FilterFW{
			Mark: 1,
		},
		Parent:  0x2,
		ToClass: 0x10,
		Prio:    10,
		Actions: []Action{},
	})

	// tc filter add
	//	dev enp0s25
	//	parent 2:
	//	prio 30
	//	protocol ip
	//	handle ::2
	//	u32
	//		ht 1:42
	//		match ip dport 25 0xffff
	//		action police rate 2Mbit burst 200K drop
	tc.AddFilter(&i, Filter{
		Type:     "u32",
		Protocol: "ip",
		Spec: &FilterU32{
			Table:     0x1,
			Bucket:    0x42,
			Index:     0x2,
			LinkTable: -1,
			Options: []Option{
				Option{Name: "match", Value: "ip src 10.30.40.200"},
			},
		},
		Parent:  0x1,
		ToClass: 0,
		Prio:    30,
		Actions: []Action{
			Action{Type: "police", Options: []string{"rate", "2Mbit", "burst", "200K", "drop"}},
		},
	})

	// tc filter add
	//	dev enp0s25
	//	parent 2:
	//	prio 20
	//	protocol ip
	//	handle ::3
	//	u32
	//		match ip dport 80 0xffff
	//		action pedit munge offset 22 u16 set 11500 pipe
	//		action csum tcp pipe
	//		action mirred egress mirror dev enp25s0
	tc.AddFilter(&i, Filter{
		Type:     "u32",
		Protocol: "ip",
		Spec: &FilterU32{
			Index:     0x3,
			LinkTable: -1,
			Options: []Option{
				Option{Name: "match", Value: "ip dst 10.30.40.0/24"},
			},
		},
		Parent:  0x2,
		ToClass: 0,
		Prio:    20,
		Actions: []Action{
			Action{Type: "pedit", Options: []string{"munge", "offset", "22", "u16", "set", "11500", "pipe"}},
			Action{Type: "csum", Options: []string{"tcp", "pipe"}},
			Action{Type: "mirred", Options: []string{"egress", "mirror", "dev", "wlp3s0"}},
		},
	})

	// tc qdisc del
	//	dev enp0s25
	//	root
	tc.DelQdisc(&i, QDisc{})

	// tc qdisc del
	//	dev enp0s25
	//	ingress
	tc.DelQdisc(&i, QDisc{Type: INGRESS_VALUE})

	fmt.Println("==== Second get")

	tc.Prepare()
	tc.Get()
	if err := tc.Commit(); err != nil {
		return err
	}
	ifaces_new = db.ifaces
	if err != nil {
		return err
	}

	var check = deep.Equal(i, ifaces_new[test_dev_in])
	for _, c := range check {
		fmt.Println(c)
	}
	return nil
}

func ApiTest(filename string) error {

	var DB = Database{Tc: &ExecTcBind{}}
	if err := DB.Load(); err != nil {
		return err
	}
	defer DB.Close()

	if configBytes, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else {
		var cfgs map[string]UserClass
		if err := yaml.Unmarshal(configBytes, &cfgs); err != nil {
			return err
		}

		for clsname, cls := range cfgs {
			if len(cls.Name) == 0 {
				cls.Name = clsname
			}
			DB.AddClass(cls)
		}
	}

	if err := DB.InitIface(test_dev_in, DIR_IN, "default"); err != nil {
		return err
	}
	if err := DB.InitIface(test_dev_out, DIR_OUT, "default"); err != nil {
		return err
	}

	var ch = make(chan error)

	go func() {
		if err := DB.NewUser(user1, group1); err != nil {
			ch <- err
			return
		}
		fmt.Println("\n====Assign user 1")
		if err := DB.AssignUser(test_dev_in, test_dev_out, user1); err != nil {
			ch <- err
			return
		}
		fmt.Println("\n====Remove user 1")
		if err := DB.DelUser(user1); err != nil {
			ch <- err
			return
		}
		ch <- nil
	}()

	go func() {
		if err := DB.NewUser(user2, group2); err != nil {
			ch <- err
			return
		}
		fmt.Println("\n====Assign user 2")
		if err := DB.AssignUser(test_dev_in, test_dev_out, user2); err != nil {
			ch <- err
			return
		}
		fmt.Println("\n====Change user 2")
		if err := DB.ChangeUser(user2, group3); err != nil {
			ch <- err
			return
		}
		ch <- nil
	}()

	go func() {
		if err := DB.NewUser(user3, group3); err != nil {
			ch <- err
			return
		}
		fmt.Println("\n====Assign user 3")
		if err := DB.AssignUser(test_dev_in, test_dev_out, user3); err != nil {
			ch <- err
			return
		}
		fmt.Println("\n====Deassign user 3")
		if err := DB.DeassignUser(test_dev_in, user3); err != nil {
			ch <- err
			return
		}
		if err := DB.DeassignUser(test_dev_out, user3); err != nil {
			ch <- err
			return
		}
		ch <- nil
	}()

	var e1 = <-ch
	var e2 = <-ch
	var e3 = <-ch

	if e1 != nil {
		return e1
	}
	if e2 != nil {
		return e2
	}
	if e3 != nil {
		return e3
	}

	return nil
}
