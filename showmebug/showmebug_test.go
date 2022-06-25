package showmebug

import (
	"fmt"
	"testing"
)

func TestDirReduce(t *testing.T) {
	params := []string{"EAST", "SOUTH", "NORTH", "WEST"}
	result := dirReduce(params)
	for _, v := range result {
		fmt.Println(v)
	}
}
