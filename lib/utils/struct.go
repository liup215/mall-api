package utils

import (
	"reflect"
	"strings"
)

func StructToMap(obj interface{}) map[string]interface{} {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	objMap := make(map[string]interface{}, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		key := f.Tag.Get("json")
		if key == "-" {
			break
		} else if key == "" {
			key = f.Name
		}

		ks := strings.Split(key, ",omitempty")
		key = ks[0]
		objMap[key] = v.FieldByName(f.Name).Interface()
	}

	return objMap
}
