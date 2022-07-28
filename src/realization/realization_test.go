package realization

import (
	"fmt"
	"testing"
)

func TestRealization(t *testing.T) {
	var Real User
	Real.SetName("王康")
	Real.SetId("2018")
	real2 := Real.GetUserData()
	fmt.Println()
}
