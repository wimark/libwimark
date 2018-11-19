package libtc

type ExecTcBind struct {
	lastError error
	ifaces    map[string]Iface
	db        *Database
}

func (self *ExecTcBind) Prepare() {
	self.lastError = nil
	self.ifaces = nil
}

func (self *ExecTcBind) Commit() error {
	// TODO batch exec
	if self.db != nil && self.ifaces != nil {
		self.db.SetInterfaces(self.ifaces)
	}
	return self.lastError
}

func (self *ExecTcBind) Error(err error) {
	self.lastError = err
}

func (self *ExecTcBind) Init(db *Database) {
	self.db = db
}
