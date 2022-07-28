package flag

import (
	"fmt"
	"testing"
)

func TestFlag(t *testing.T) {
	user := ArgsInitUser()
	fmt.Println(user)
}
