package libtc

import (
	"errors"
	"fmt"
	"strconv"
)

func (self *ExecTcBind) DelFilter(dev *Iface, f Filter) {
	// Example:
	//	tc filter del
	//		dev wlp3s0 parent 2: prio 20 handle ::3 u32

	if self.lastError != nil {
		return
	}

	// TODO checks
	var qdisc, ok = dev.Discs[f.Parent]
	if !ok {
		self.lastError = errors.New("No qdisc in filter del")
		return
	}

	// delete
	var params = []string{
		FILTER_OBJECT, DEL_COMMAND,
		DEV_PARAM, dev.Name,
		PARENT_PARAM, fmt.Sprintf("%x:", f.Parent),
		PRIO_PARAM, strconv.Itoa(f.Prio),
	}
	params = append(params, f.Spec.MakeDelParams()...)

	if _, err := execute(TC_EXECUTABLE, params...); err != nil {
		self.lastError = err
		return
	}
	for i, filt := range qdisc.Filters {
		if filt.Spec.Handle() == f.Spec.Handle() {
			qdisc.Filters = append(qdisc.Filters[:i], qdisc.Filters[i+1:]...)
			break
		}
	}
	dev.Discs[f.Parent] = qdisc
}

func (self *ExecTcBind) DelClass(dev *Iface, c Class) {
	// Example:
	//	tc class del
	//		dev eth0 classid 1:1

	if self.lastError != nil {
		return
	}

	// TODO checks
	var qdisc, ok = dev.Discs[c.ParentDisc]
	if !ok {
		self.lastError = errors.New("No qdisc in class add")
		return
	}

	// update
	var params = []string{
		CLASS_OBJECT, DEL_COMMAND,
		DEV_PARAM, dev.Name,
		CLASSID_PARAM, fmt.Sprintf("%x:%x", c.ParentDisc, c.Handle),
	}

	if _, err := execute(TC_EXECUTABLE, params...); err != nil {
		self.lastError = err
		return
	}
	delete(qdisc.Classes, c.Handle)
	dev.Discs[c.ParentDisc] = qdisc
}

func (self *ExecTcBind) DelQdisc(dev *Iface, qdisc QDisc) {
	// Example:
	//	tc qdics del
	//		dev eth0 parent 1:10 handle 2:

	if self.lastError != nil {
		return
	}

	// TODO checks

	// delete
	var params = []string{
		QDISC_OBJECT, DEL_COMMAND,
		DEV_PARAM, dev.Name,
	}
	if qdisc.Type == INGRESS_VALUE {
		params = append(params, INGRESS_VALUE)
	} else if qdisc.ParentClass == 0 {
		params = append(params, ROOT_VALUE)
	} else {
		params = append(params,
			PARENT_PARAM, fmt.Sprintf("%x:%x", qdisc.ParentDisc, qdisc.ParentClass),
			HANDLE_PARAM, fmt.Sprintf("%x:", qdisc.Handle),
		)
	}

	if _, err := execute(TC_EXECUTABLE, params...); err != nil {
		self.lastError = err
		return
	}
	delete(dev.Discs, qdisc.Handle)
	for key, d := range dev.Discs {
		if d.ParentDisc == qdisc.Handle {
			delete(dev.Discs, key)
		}
	}
}
