package libtc

import (
	"fmt"
	"regexp"
	"strings"
)

func (self *ExecTcBind) Get() {

	if self.lastError != nil {
		return
	}

	var params = []string{
		LINK_OBJECT,
		SHOW_COMMAND,
	}
	self.ifaces = map[string]Iface{}
	var out, err = execute(IP_EXECUTABLE, params...)
	if err != nil {
		self.lastError = err
		return
	}
	var res = map[string]Iface{}
	var re = regexp.MustCompile("^[[:digit:]]+:[[:space:]]+([^:]+):")
	for _, line := range out {
		var subs = re.FindStringSubmatch(line)
		if len(subs) > 1 {
			var ifname = strings.Split(subs[1], "@")[0]
			var iface, err = self.getDiscs(ifname)
			if err != nil {
				fmt.Println("  Error", err.Error())
				continue
			}
			res[ifname] = iface
		}
	}
	self.ifaces = res
}

func (self *ExecTcBind) getDiscs(iface string) (Iface, error) {

	var params = []string{
		QDISC_OBJECT,
		SHOW_COMMAND,
		DEV_PARAM,
		iface,
	}
	var out, err = execute(TC_EXECUTABLE, params...)
	if err != nil {
		return Iface{}, err
	}
	var res = Iface{
		Name:  iface,
		Discs: map[int]QDisc{},
	}
	for _, line := range out {
		var params = split(line)
		if len(params) < 4 {
			continue
		}
		if params[0] != QDISC_OBJECT {
			continue
		}
		var qdisc = QDisc{
			Type:   params[1],
			Handle: findString(params[2], 0, 16),
		}
		var optionsFrom = 4
		if params[3] == PARENT_PARAM {
			qdisc.ParentDisc = findString(params[4], 0, 16)
			qdisc.ParentClass = findString(params[4], 1, 16)
			optionsFrom = 5
		}
		qdisc.Options = parseOptions(params[optionsFrom:])
		if qdisc.Handle != 0 || (qdisc.ParentClass == 0 && qdisc.ParentDisc == 0) {

			if cls, err := self.getClasses(iface, qdisc.Handle); err == nil {
				qdisc.Classes = cls
			}
			if flt, err := self.getFilters(iface, qdisc.Handle); err == nil {
				qdisc.Filters = flt
			}
		}
		if qdisc.Handle == 0 {
			res.DefaultDiscs = append(res.DefaultDiscs, qdisc)
		} else {
			res.Discs[qdisc.Handle] = qdisc
		}
	}

	return res, nil
}

func (self *ExecTcBind) getClasses(iface string, disc int) (map[int]Class, error) {

	var params = []string{
		CLASS_OBJECT,
		SHOW_COMMAND,
		DEV_PARAM,
		iface,
		PARENT_PARAM,
		fmt.Sprintf("%x:", disc),
	}
	var out, err = execute(TC_EXECUTABLE, params...)
	if err != nil {
		return nil, err
	}
	var res = map[int]Class{}
	for _, line := range out {
		var params = split(line)
		if len(params) < 4 {
			continue
		}
		if params[0] != CLASS_OBJECT {
			continue
		}
		var cls = Class{
			Handle:     findString(params[2], 1, 16),
			ParentDisc: disc,
		}
		var optionsFrom = 4
		if params[3] == PARENT_PARAM {
			cls.ParentClass = findString(params[4], 1, 16)
			optionsFrom = 5
		}
		if optionsFrom < len(params) && params[optionsFrom] == LEAF_PARAM {
			cls.LeafDisc = findString(params[optionsFrom+1], 0, 16)
			optionsFrom += 2
		}
		cls.Options = parseOptions(params[optionsFrom:])
		res[cls.Handle] = cls
	}

	return res, nil
}

func (self *ExecTcBind) getFilters(iface string, disc int) ([]Filter, error) {

	var params = []string{
		FILTER_OBJECT,
		SHOW_COMMAND,
		DEV_PARAM,
		iface,
		PARENT_PARAM,
		fmt.Sprintf("%x:", disc),
	}
	var out, err = execute(TC_EXECUTABLE, params...)
	if err != nil {
		return nil, err
	}
	var res []Filter
	var flt Filter
	var spec []string
	for _, line := range out {
		var params = split(line)
		if len(params) == 0 {
			continue
		}
		if params[0] == FILTER_OBJECT {
			if flt.Spec != nil && flt.Spec.ParseParams(spec, &flt) == nil {
				res = append(res, flt)
			}
			flt = Filter{Parent: disc}
			spec = []string{}

			var specfrom int
			for n := 1; n < len(params); n++ {
				switch params[n] {
				case PROTO_PARAM:
					flt.Protocol = params[n+1]
				case PRIO_PARAM:
					flt.Prio = s2i(params[n+1])
					flt.Type = params[n+2]
					specfrom = n + 3
				}
				if specfrom != 0 {
					break
				}
			}
			switch flt.Type {
			case FILTER_U32:
				flt.Spec = &FilterU32{}
			case FILTER_FW:
				flt.Spec = &FilterFW{}
			}
			spec = append(spec, strings.Join(params[specfrom:], " "))

		} else {
			spec = append(spec, line)
		}
	}
	if flt.Spec != nil && flt.Spec.ParseParams(spec, &flt) == nil {
		res = append(res, flt)
	}

	return res, nil
}

func (self *ExecTcBind) GetStats(iface *Iface) (map[uint32]ClassStat, error) {
	if iface == nil {
		return nil, fmt.Errorf("Nil iface")
	}

	var params = []string{
		"-s", // with stats
		CLASS_OBJECT,
		SHOW_COMMAND,
		DEV_PARAM,
		iface.Name,
	}
	var out, err = execute(TC_EXECUTABLE, params...)
	if err != nil {
		return nil, err
	}

	var res = map[uint32]ClassStat{}
	for index, line := range out {
		var params = split(line)
		if len(params) < 4 {
			continue
		}
		if params[0] != CLASS_OBJECT {
			continue
		}

		rt := findString(params[2], 1, 16)
		lt := findString(params[2], 0, 16)

		var cls = ClassStat{
			Handle: uint32(lt)<<16 + uint32(rt),
		}
		var optionsFrom = 4
		if params[3] == PARENT_PARAM {
			cls.ParentClass = findString(params[4], 1, 16)
			optionsFrom = 5
		}
		if optionsFrom < len(params) && params[optionsFrom] == LEAF_PARAM {
			continue
			// cls.LeafDisc = findString(params[optionsFrom+1], 0, 16)
		}

		if len(out) < index+1 {
			return res, nil
		}

		line = out[index+1]
		params = split(line)
		if len(params) < 4 {
			continue
		}
		if params[0] != SENT_PARAM {
			continue
		}
		cls.Bytes = s2i(params[1])
		cls.Packets = s2i(params[3])

		// if k, ok := res[cls.Handle]; ok {
		// 	k.Bytes += cls.Bytes
		// 	k.Packets += cls.Packets
		// 	res[cls.Handle] = k
		// } else {
		res[cls.Handle] = cls
		// }

	}

	return res, nil
}
