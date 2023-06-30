package datatable

import (
	"reflect"
	"sync"
)

var (
	tagkey = "column"
	mux    sync.Mutex
)

func SetTagkey(key string) {
	mux.Lock()
	defer mux.Unlock()
	tagkey = key
}

func GetTagkey() string {
	mux.Lock()
	defer mux.Unlock()
	return tagkey
}

func tagValues(data interface{}) []string {
	var v = reflect.ValueOf(data)
	if v.Len() > 0 {
		vt := v.Index(0).Type()
	
		cols := []string{}
		for i := 0; i < vt.NumField(); i++ {
			tv := vt.Field(i).Tag.Get(tagkey)
			if len(tv) > 0 && tv != "-" {
				cols = append(cols, tv)
			}
		}

		return cols
	}

	return []string{"NO DATA"}
}

// getFieldValue は、タグのキーを指定してフィールドの値を抽出します。
func getFieldValue(s interface{}, i int) interface{} {
	v := reflect.ValueOf(s)
	vt := v.Type()

	var cnt int
	for j := 0; j < vt.NumField(); j++ {
		tv := vt.Field(j).Tag.Get(tagkey)
		if len(tv) > 0 && tv != "-" {
			if cnt == i {
				return v.Field(j).Interface()
			}
			cnt++
		}
	}
	
	return nil
}
