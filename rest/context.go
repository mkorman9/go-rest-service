package rest

import (
	"sync"
	"strings"
)

var RestConfig = "restConfig"

var instance *AppContext
var lock sync.Once

type AppContext struct {
	members map[string]interface{}
}

func (context *AppContext) GetMember(name string) interface{} {
	return context.members[name]
}

func (context *AppContext) GetDatasourcesList() []string {
	datasources := make([]string, 0)
	for key, _ := range context.members {
		if strings.HasPrefix(key, "db_") {
			datasources = append(datasources, key)
		}
	}

	return datasources
}

func (context *AppContext) AddMember(name string, value interface{}) {
	context.members[name] = value
}

func GetContext() *AppContext {
	lock.Do(func() {
		instance = &AppContext{}
		instance.members = make(map[string]interface{})
	})
	return instance
}
