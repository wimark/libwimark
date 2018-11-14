package libtc

import (
	"errors"
	"fmt"
	"strconv"
)

func AddFilter(dev *Iface, f Filter) error {
	// Example:
	//	tc filter add
	//		dev wlp3s0 parent 2: prio 20 protocol ip handle ::3 u32
	//		match ip dport 80 0xffff
	//		action pedit munge offset 22 u16 set 11500 pipe
	//		action csum tcp pipe
	//		action mirred egress mirror dev enp0s25

	// TODO checks
	var qdisc, ok = dev.Discs[f.Parent]
	if !ok {
		return errors.New("No qdisc in filter add")
	}

	// add
	var params = []string{
		FILTER_OBJECT, ADD_COMMAND,
		DEV_PARAM, dev.Name,
		PARENT_PARAM, fmt.Sprintf("%x:", f.Parent),
		PRIO_PARAM, strconv.Itoa(f.Prio),
		PROTO_PARAM, f.Protocol,
	}
	var opts = f.Spec.MakeAddParams()
	var acts = act2params(f.Actions)
	params = append(params, opts...)
	if f.ToClass != 0 {
		params = append(params, CLASSID_PARAM, fmt.Sprintf("%x:%x", f.Parent, f.ToClass))
	}
	params = append(params, acts...)

	var _, err = execute(TC_EXECUTABLE, params...)
	if err == nil {
		qdisc.Filters = append(qdisc.Filters, f)
		dev.Discs[f.Parent] = qdisc
	}
	return err
}

func AddClass(dev *Iface, c Class) error {
	// Example:
	//	tc class add
	//		dev eth0 parent 1: classid 1:1 htb
	//		rate 90kbps ceil 90kbps

	// TODO checks
	var qdisc, ok = dev.Discs[c.ParentDisc]
	if !ok {
		return errors.New("No qdisc in class add")
	}

	// add
	var params = []string{
		CLASS_OBJECT, ADD_COMMAND,
		DEV_PARAM, dev.Name,
		PARENT_PARAM, fmt.Sprintf("%x:%x", c.ParentDisc, c.ParentClass),
		CLASSID_PARAM, fmt.Sprintf("%x:%x", c.ParentDisc, c.Handle),
		qdisc.Type,
	}
	var opts = opt2params(c.Options)
	params = append(params, opts...)

	var _, err = execute(TC_EXECUTABLE, params...)
	if err == nil {
		if qdisc.Classes == nil {
			qdisc.Classes = map[int]Class{}
		}
		qdisc.Classes[c.Handle] = c
		dev.Discs[c.ParentDisc] = qdisc
	}
	return err
}

func AddQdisc(dev *Iface, qdisc QDisc) error {
	// Example:
	//	tc qdics add
	//		dev eth0 parent 1:10 handle 2: htb
	//		default 20

	// TODO checks

	// add
	var params = []string{
		QDISC_OBJECT, ADD_COMMAND,
		DEV_PARAM, dev.Name,
	}
	if qdisc.Type == INGRESS_VALUE {
		params = append(params, INGRESS_VALUE)
		qdisc.Handle = 0xffff
		qdisc.ParentClass = 0xfff1
		qdisc.ParentDisc = 0xffff
	} else {
		if qdisc.ParentClass == 0 {
			params = append(params, ROOT_VALUE)
		} else {
			params = append(params,
				PARENT_PARAM, fmt.Sprintf("%x:%x", qdisc.ParentDisc, qdisc.ParentClass),
			)
		}
		params = append(params,
			HANDLE_PARAM, fmt.Sprintf("%x:", qdisc.Handle),
			qdisc.Type,
		)
		var opts = opt2params(qdisc.Options)
		params = append(params, opts...)
	}

	var _, err = execute(TC_EXECUTABLE, params...)
	if err == nil {
		dev.DefaultDiscs = nil
		dev.Discs[qdisc.Handle] = qdisc
		if qdisc.ParentDisc != 0 {
			var pdisc = dev.Discs[qdisc.ParentDisc]
			var pcls = pdisc.Classes[qdisc.ParentClass]
			pcls.LeafDisc = qdisc.Handle
			if pdisc.Classes == nil {
				pdisc.Classes = map[int]Class{}
			}
			pdisc.Classes[qdisc.ParentClass] = pcls
			dev.Discs[qdisc.ParentDisc] = pdisc
		}
	}
	return err
}
