package libwimark

import (
	"sync"
)

// MACPrefixVendorMap is singleton map that contains MAC prefix-vendor pairs
var MACPrefixVendorMap ManufMap

// MACPrefixVendor тип данных mac префикс и производитель оборудования
type MACPrefixVendor struct {
	Prefix string `bson:"prefix"`
	Vendor string `bson:"vendor"`
}

// ManufMap struct for store MAC prefix-vendor pairs
type ManufMap struct {
	data   map[string]string
	m      sync.RWMutex
	inited bool
}

//GenerateMap принмает массив, переводит его в мапу
func (e *ManufMap) GenerateMap(data []MACPrefixVendor) {
	e.data = make(map[string]string, len(data))
	for _, v := range data {
		e.Add(v.Prefix, v.Vendor)
	}
	e.inited = true
}

func (e *ManufMap) check() {
	if e.data == nil {
		e.data = make(map[string]string)
		e.inited = false
	}
}

//Len возвращает длинну
func (e *ManufMap) Len() int {
	e.m.RLock()
	e.check()
	defer e.m.RUnlock()
	return len(e.data)
}

//Add добавляет канал в базу
func (e *ManufMap) Add(key, c string) {
	e.m.Lock()
	e.check()
	defer e.m.Unlock()
	e.data[key] = c
}

//Get возвращает значение из мапы
func (e *ManufMap) Get(key string) string {
	e.m.RLock()
	e.check()
	defer e.m.RUnlock()
	return e.data[key]
}

//Delete удаляет значение из базы
func (e *ManufMap) Delete(key string) {
	e.m.Lock()
	e.check()
	defer e.m.Unlock()
	delete(e.data, key)
}
