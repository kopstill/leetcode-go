package showmebug

import (
	"fmt"
	"testing"
)

func TestDirReduce(t *testing.T) {
	// params := []string{"WEST", "SOUTH", "EAST", "NORTH"}
	// params := []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
	// params := []string{"NORTH", "SOUTH", "EAST", "WEST"}
	// params := []string{"NORTH", "EAST", "WEST", "SOUTH", "WEST", "WEST"}
	// params := []string{"NORTH", "WEST", "SOUTH", "EAST", "WEST"}
	params := []string{"NORTH", "WEST", "SOUTH", "EAST", "WEST", "NORTH", "EAST", "XXX", "SOUTH", "WEST", "WEST", "EAST", "EAST", "XXX"}
	result := dirReduce(params)
	for _, v := range result {
		fmt.Println(v)
	}
}
