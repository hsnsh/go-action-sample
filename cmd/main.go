//package main
//
//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	fmt.Println("Welcome Action Sample Application")
//}

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	type MyEntity struct {
		myColumn string `gorm:"column:name;not null;size:250;"`
	}

	isRequired, maxLength := GetGormTagDetails(MyEntity{}, "myColumn")

	fmt.Printf("Type Name : %s\n", reflect.TypeOf(MyEntity{}).Name())
	fmt.Printf("Type Field Name : %s\n", "myColumn")
	fmt.Printf("Type Field IsRequired : %v\n", isRequired)
	fmt.Printf("Type Field MaxLength : %v\n", maxLength)

}

func GetGormTagDetails(model interface{}, fieldName string) (isRequired bool, maxLength int) {

	if field, exist := reflect.TypeOf(model).FieldByName(fieldName); exist {
		rules := strings.Split(field.Tag.Get("gorm"), ";")
		for _, r := range rules {
			if len(r) > 0 {
				//fmt.Printf("Rule -> %s\n", r)
				if r == "not null" {
					isRequired = true
				}
				if strings.Contains(r, "size") {
					maxLength, _ = strconv.Atoi(string(strings.Split(r, ":")[1]))
				}
			}
		}

	} else {
		fmt.Printf("Type Field Name : %s not found\n", fieldName)
	}

	return isRequired, maxLength
}
