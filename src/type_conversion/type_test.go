package type_conversion

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	data := "231"
	var reda int
	fmt.Printf("the type is %T \n", reda)
	GetValue(data, &reda)
	fmt.Printf("the type is %T", reda)
}
