package conv

import (
	"errors"
	"reflect"
	"strings"
)

const (
	OptIgnore    = "-"         // 忽略该字段
	OptOmitempty = "omitempty" // 该字段为空时忽略该字段
	OptDive      = "dive"      // 递归遍历该字段
)

// 结构体字段的标签选项
const (
	flagIgnore = 1 << iota
	flagOmiEmpty
	flagDive
)

// readTag 读取结构体字段的标签信息
func readTag(f reflect.StructField, tag string) (string, int) {
	val, ok := f.Tag.Lookup(tag)
	fieldTag := ""
	flag := 0
	if !ok {
		return f.Name, flag
	}
	opts := strings.Split(val, ",")
	fieldTag = opts[0]
	for i := 1; i < len(opts); i++ {
		switch opts[i] {
		case OptIgnore:
			flag |= flagIgnore
		case OptOmitempty:
			flag |= flagOmiEmpty
		case OptDive:
			flag |= flagDive
		}
	}
	return fieldTag, flag
}

func Struct2Map(s interface{}) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("input is not a struct")
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.FieldByName(field.Name)
		tagVal, flag := readTag(field, "map")
		if flag&flagIgnore != 0 {
			continue
		}
		if fieldValue.IsZero() && flag&flagOmiEmpty != 0 {
			continue
		}
		if fieldValue.Kind() == reflect.Struct {
			if flag&flagDive != 0 {
				innerMap, err := Struct2Map(fieldValue.Interface())
				if err != nil {
					return nil, err
				}
				for k, v := range innerMap {
					res[k] = v
				}
			} else {
				res[tagVal] = fieldValue.Interface()
			}
		} else {
			res[tagVal] = fieldValue.Interface()
		}
	}
	return res, nil
}
