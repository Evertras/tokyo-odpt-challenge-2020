package esdata

import "strings"

func RemoveType(raw string) string {
	split := strings.SplitN(raw, ":", 2)

	if len(split) == 1 {
		return raw
	}

	return split[1]
}
