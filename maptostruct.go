package main

//test.....
import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
	//	"unsafe"
)
//2222222222222
type User struct {
	Username string            `json:"user"`
	Age      string            `json:"age"`
	Times    time.Time         `json:"t"`
	S        map[string]string `json:"s"`
	TT       string            `json:"tt"`
}

func main() {
	m := map[string]interface{}{"age": 45, "user": "hangsan", "t": "2012-12-12 12:12:12", "s": map[string]string{"g": "h"}}
	u := &User{}
	err := MapToStruct(m, u)
	fmt.Println("....", err)
	fmt.Printf("%+v", u)

}

//支持基本类型的转换
func MapToStruct(m map[string]interface{}, u interface{}) error {
	map_lenth := len(m)
	t := reflect.TypeOf(u).Elem()

	for i := 0; i < t.NumField(); i++ {
		structtype := t.FieldByIndex([]int{i})
		tag := structtype.Tag.Get("json")
		mv, ok := m[tag]

		if !ok {
			fmt.Printf("[maptostruct] map has not the key:%s\n", tag)
			continue
		}

		structvalue := reflect.ValueOf(u).Elem()

		structfieldvalue := structvalue.FieldByName(structtype.Name)

		if !structfieldvalue.CanSet() {

			return errors.New(fmt.Sprintf("Cannot set %s field value\n", structtype.Name))

		}

		if structtype.Type == reflect.TypeOf(mv) {
			structfieldvalue.Set(reflect.ValueOf(mv))
			map_lenth -= 1
			continue
		}

		v, err := TypeConversion(structtype.Type.String(), fmt.Sprintf("%v", mv))
		if err != nil {
			return err

		}
		structfieldvalue.Set(v)
		map_lenth -= 1

	}

	if map_lenth > 0 {
		return errors.New("map not cnversion compelety")
	}
	return nil

}

func TypeConversion(stype string, value string) (reflect.Value, error) {
	if stype == "string" {
		return reflect.ValueOf(value), nil

	}

	switch stype {
	case "int":
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err

	case "int64":
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err

	case "float64":
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	case "time.Time":
		i, err := time.Parse("2006-01-02 15:04:05", value)
		return reflect.ValueOf(i), err
	default:
		return reflect.ValueOf(value), errors.New(fmt.Sprintf("unkonw type:%s", stype))
	}
	// []int,map[string]string

}
