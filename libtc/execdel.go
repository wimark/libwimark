package libtc

import (
	"errors"
	"fmt"
	"strconv"
)

func DelFilter(dev *Iface, f Filter) error {
	// Example:
	//	tc filter del
	//		dev wlp3s0 parent 2: prio 20 handle ::3 u32

	// TODO checks
	var qdisc, ok = dev.Discs[f.Parent]
	if !ok {
		return errors.New("No qdisc in filter del")
	}

	// delete
	var params = []string{
		FILTER_OBJECT, DEL_COMMAND,
		DEV_PARAM, dev.Name,
		PARENT_PARAM, fmt.Sprintf("%x:", f.Parent),
		PRIO_PARAM, strconv.Itoa(f.Prio),
	}
	params = append(params, f.Spec.MakeDelParams()...)

	var _, err = execute(TC_EXECUTABLE, params...)
	if err == nil {
		for i, filt := range qdisc.Filters {
			if filt.Spec.Handle() == f.Spec.Handle() {
				qdisc.Filters = append(qdisc.Filters[:i], qdisc.Filters[i+1:]...)
				break
			}
		}
		dev.Discs[f.Parent] = qdisc
	}
	return err
}

func DelClass(dev *Iface, c Class) error {
	// Example:
	//	tc class del
	//		dev eth0 classid 1:1

	// TODO checks
	var qdisc, ok = dev.Discs[c.ParentDisc]
	if !ok {
		return errors.New("No qdisc in class add")
	}

	// update
	var params = []string{
		CLASS_OBJECT, DEL_COMMAND,
		DEV_PARAM, dev.Name,
		CLASSID_PARAM, fmt.Sprintf("%x:%x", c.ParentDisc, c.Handle),
	}

	var _, err = execute(TC_EXECUTABLE, params...)
	if err == nil {
		delete(qdisc.Classes, c.Handle)
		dev.Discs[c.ParentDisc] = qdisc
	}
	return err
}

func DelQdisc(dev *Iface, qdisc QDisc) error {
	// Example:
	//	tc qdics del
	//		dev eth0 parent 1:10 handle 2:

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

	var _, err = execute(TC_EXECUTABLE, params...)
	if err == nil {
		delete(dev.Discs, qdisc.Handle)
		for key, d := range dev.Discs {
			if d.ParentDisc == qdisc.Handle {
				delete(dev.Discs, key)
			}
		}
	}
	return err
}
