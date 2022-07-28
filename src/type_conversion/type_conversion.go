package type_conversion

import (
	"fmt"
	"strconv"
)

func GetValue(k string, val interface{}) {
	*(val.(*int)), _ = strconv.Atoi(k)
	fmt.Printf("the type is %T \n", *(val.(*int)))
}
