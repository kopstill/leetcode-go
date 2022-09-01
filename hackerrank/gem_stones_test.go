package hackerrank

import (
	"fmt"
	"testing"
)

func TestGemstones(t *testing.T) {
	// rocks := []string{"abcdde"}
	// rocks := []string{"abcdde", "baccd"}
	// rocks := []string{"abcdde", "baccd", "eeabg"}
	rocks := []string{"abcdde", "baccd", "eeabg", "bcdeg"}
	fmt.Println(gemstones(rocks))
}
