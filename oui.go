package libwimark

import (
	"strconv"
	"sync"
)

// MACPrefixVendorMap is singleton map that contains MAC prefix-vendor pairs
var MACPrefixVendorMap ManufMap

// MACPrefixVendor тип данных mac префикс и производитель оборудования
type MACPrefixVendor struct {
	Prefix string `bson:"_id" json:"prefix"`
	Vendor string `bson:"vendor" json:"vendor"`
}

// ManufMap struct for store MAC prefix-vendor pairs
type ManufMap struct {
	data   map[int64]string
	m      sync.RWMutex
	inited bool
}

//GenerateMap принмает массив, переводит его в мапу
func (e *ManufMap) GenerateMap(data []MACPrefixVendor) {
	e.data = make(map[int64]string, len(data))
	for _, v := range data {
		e.Add(v.Prefix, v.Vendor)
	}
	e.inited = true
}

func (e *ManufMap) check() {
	if e.data == nil {
		e.data = make(map[int64]string)
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
	v, err := strconv.ParseInt(key, 16, 64)
	if err == nil {
		e.data[v] = c
	}
}

//Get возвращает значение из мапы
func (e *ManufMap) Get(key string) string {
	e.m.RLock()
	e.check()
	defer e.m.RUnlock()
	v, err := strconv.ParseInt(key, 16, 64)
	if err == nil {
		return e.data[v]
	}
	return ""
}

//Delete удаляет значение из базы
func (e *ManufMap) Delete(key string) {
	e.m.Lock()
	e.check()
	defer e.m.Unlock()
	v, err := strconv.ParseInt(key, 16, 64)
	if err == nil {
		delete(e.data, v)
	}
}

//Add добавляет канал в базу
func (e *ManufMap) AddI(key int64, c string) {
	e.m.Lock()
	e.check()
	defer e.m.Unlock()
	e.data[key] = c
}

//Get возвращает значение из мапы
func (e *ManufMap) GetI(key int64) string {
	e.m.RLock()
	e.check()
	defer e.m.RUnlock()

	return e.data[key]

}

//Delete удаляет значение из базы
func (e *ManufMap) DeleteI(key int64) {
	e.m.Lock()
	e.check()
	defer e.m.Unlock()
	delete(e.data, key)

}
