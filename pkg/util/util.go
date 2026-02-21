package util

import "strings"

// StripExtension removes the file extension from a filename
func StripExtension(name string) string {
	if dot := strings.LastIndex(name, "."); dot != -1 {
		return name[:dot]
	}
	return name
}
