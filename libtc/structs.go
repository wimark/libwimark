package libtc

type Option struct {
	Name  string
	Value string
}

type Action struct {
	Type    string
	Order   int
	Options []string
}

type FilterSpec interface {
	Handle() string
	Id() int
	MakeAddParams() []string
	MakeDelParams() []string
	ParseParams([]string, *Filter) error
}

type Filter struct {
	Type     string
	Protocol string
	Spec     FilterSpec
	Parent   int
	ToClass  int
	Prio     int
	Actions  []Action
}

type Class struct {
	Handle      int
	ParentDisc  int
	ParentClass int
	LeafDisc    int
	Options     []Option
}

type QDisc struct {
	Type        string
	ParentDisc  int
	ParentClass int
	Handle      int
	Options     []Option
	Classes     map[int]Class
	Filters     []Filter
}

type Iface struct {
	Name         string
	Discs        map[int]QDisc
	DefaultDiscs []QDisc
}

type ClassStat struct {
	Handle      uint32
	ParentDisc  int
	ParentClass int
	// LeafDisc    int
	Bytes   int
	Packets int
}
