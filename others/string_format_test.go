package others

import (
	"fmt"
	"testing"
)

func TestFormatString(t *testing.T) {
	// format := "{{aaa}}{{bbb}}"
	// format := "{{{{aaa}}{{{{bbb}}"
	// format := "{{}{{{aaa}}{}{{{bbb}}}{{aaa}}{{}{{aaa}}}}{{}{{bbb}}"
	format := " {{{           aaa    }} {{   bbb     } }. "
	args := map[string]string{
		"aaa": "AAA",
		"bbb": "BBB",
	}
	fmt.Println(formatString(format, args))
}
