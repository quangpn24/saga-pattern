package utils

import "strings"

func ContainFirst(elems []string, v string) bool {
	for _, s := range elems {
		if strings.HasPrefix(v, s) {
			return true
		}
	}

	return false
}
