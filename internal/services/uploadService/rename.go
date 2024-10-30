package uploadService

import "strings"

func Rename(name string, id string) string {
	names := strings.Split(name, ".")
	if len(name) < 1 {
		return ""
	}
	return id + "." + names[1]
}
