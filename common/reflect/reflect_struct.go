package reflect

import (
	"errors"
	"fmt"
	"reflect"
)

func ParseStruct(v interface{}, out interface{}) (err error) {
	refValue := reflect.ValueOf(v) // value
	refType := reflect.TypeOf(v)   // type
	kind := refValue.Kind()        // basic type

	refValueOut := reflect.ValueOf(out) // value
	refTypeOut := reflect.TypeOf(out)   // type
	kindOut := refTypeOut.Kind()        // basic type

	fmt.Println(" Kind:", kind)
	if kind != reflect.Struct || kindOut != reflect.Struct {
		err = errors.New("Parse value is not struct")
		return
	}

	fieldCount := refValue.NumField()       // field count
	fieldCountOut := refValueOut.NumField() // field count
	fmt.Println("fieldCount:", fieldCount)
	fmt.Println("fieldCountOut:", fieldCountOut)
	fmt.Println("")

	for i := 0; i < fieldCount; i++ {
		fieldType := refType.Field(i)   // field type
		fieldValue := refValue.Field(i) // field vlaue

		for j := 0; j < fieldCountOut; j++ {
			fieldTypeOut := refTypeOut.Field(j) // field type
			// fieldValueOut := refValueOut.Field(j) // field vlaue
			if fieldType.Name == fieldTypeOut.Name {
				fmt.Println("Find fieldType:", fieldType.Name)
				fmt.Println("fieldValue:", fieldValue)
				//需要实现赋值
				// fieldValueOut.Set(fieldValue.Elem())
			}
		}
		fmt.Println("")
	}
	return
}
