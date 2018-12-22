package cache

import (
	"time"
)

const timeSpan = 1800

var cacheMap = make(map[string]Updater)

type Updater interface {
	UpdateTime() int
	Name() string
}

type Generator func(city string) Updater

func GetOrGenerate(city string, generate Generator) interface{} {
	su, ok := cacheMap[city]
	if !ok || int(time.Now().Unix())-su.UpdateTime() > timeSpan {
		su = generate(city)
		if su != nil {
			if su.Name() == "" {
				delete(cacheMap, city)
			} else {
				cacheMap[city] = su
			}
		}
	}
	return su
}
