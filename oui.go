package libwimark

import (
	"sync"
)

//ManufData тип данных mac префикс и производитель оборудования
type ManufData struct {
	Prefix string `bson:"prefix"`
	Vendor string `bson:"vendor"`
}

type ManufMap struct {
	isInit bool
	data   map[string]string
	m      sync.Mutex
}

//GenerateMap принмает массив, переводит его в мапу
func (e *ManufMap) GenerateMap(data []ManufData) {
	e.chek()
	for _, v := range data {
		e.Add(v.Prefix, v.Vendor)
	}
}

func (e *ManufMap) chek() {
	if !e.isInit {
		e.Init()
		e.isInit = true
	}
}

//Init инициализирует мапу
func (e *ManufMap) Init() {
	e.data = make(map[string]string)
	e.isInit = true
}

//GetLen возвращает длинну
func (e ManufMap) GetLen() int {
	e.m.Lock()
	e.chek()
	defer e.m.Unlock()
	return len(e.data)
}

//Add добавляет канал в базу
func (e *ManufMap) Add(key string, c string) {
	e.m.Lock()
	e.chek()
	defer e.m.Unlock()
	e.data[key] = c
}

//Get возвращает значение из мапы
func (e *ManufMap) Get(key string) string {
	e.m.Lock()
	e.chek()
	defer e.m.Unlock()
	return e.data[key]
}

//Delete удаляет значение из базы
func (e *ManufMap) Delete(key string) {
	e.m.Lock()
	e.chek()
	defer e.m.Unlock()
	delete(e.data, key)
}
