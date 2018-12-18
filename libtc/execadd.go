package libtc

import (
	"errors"
	"fmt"
	"strconv"
)

func (self *ExecTcBind) AddFilter(dev *Iface, f Filter) {
	// Example:
	//	tc filter add
	//		dev wlp3s0 parent 2: prio 20 protocol ip handle ::3 u32
	//		match ip dport 80 0xffff
	//		action pedit munge offset 22 u16 set 11500 pipe
	//		action csum tcp pipe
	//		action mirred egress mirror dev enp0s25

	if self.lastError != nil {
		return
	}

	// TODO checks
	var qdisc, ok = dev.Discs[f.Parent]
	if !ok {
		self.lastError = errors.New("No qdisc in filter add")
		return
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

	if _, err := execute(TC_EXECUTABLE, params...); err != nil {
		self.lastError = err
		return
	}
	qdisc.Filters = append(qdisc.Filters, f)
	dev.Discs[f.Parent] = qdisc
}

func (self *ExecTcBind) AddClass(dev *Iface, c Class) {
	// Example:
	//	tc class add
	//		dev eth0 parent 1: classid 1:1 htb
	//		rate 90kbps ceil 90kbps

	if self.lastError != nil {
		return
	}

	// TODO checks
	var qdisc, ok = dev.Discs[c.ParentDisc]
	if !ok {
		self.lastError = errors.New("No qdisc in class add")
		return
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

	if _, err := execute(TC_EXECUTABLE, params...); err != nil {
		self.lastError = err
		return
	}
	if qdisc.Classes == nil {
		qdisc.Classes = map[int]Class{}
	}
	qdisc.Classes[c.Handle] = c
	dev.Discs[c.ParentDisc] = qdisc
}

func (self *ExecTcBind) AddQdisc(dev *Iface, qdisc QDisc) {
	// Example:
	//	tc qdics add
	//		dev eth0 parent 1:10 handle 2: htb
	//		default 20

	if self.lastError != nil {
		return
	}

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

	if _, err := execute(TC_EXECUTABLE, params...); err != nil {
		self.lastError = err
		return
	}
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
