package others

import (
	"strings"
)

func formatString(format string, args map[string]string) string {
	if len(strings.TrimSpace(format)) == 0 || args == nil {
		panic("Required parameters missing.")
	}

	left, right := -1, -1
	for i := 0; i < len(format)-1; i++ {
		if format[i] == '{' && format[i+1] == '{' {
			left = i
		}
		if format[i] == '}' && format[i+1] == '}' {
			right = i
		}

		if left != -1 && right != -1 && left < right {
			key := strings.TrimSpace(format[left+2 : right])
			if key == "" {
				panic("Match key must not be empty.")
			}
			value, exists := args[key]
			if !exists {
				panic("Match value must not be empty.")
			}

			format = format[0:left] + value + format[right+2:]

			i = left + len(value) - 1
			left = -1
			right = -1
		}
	}

	return format
}
