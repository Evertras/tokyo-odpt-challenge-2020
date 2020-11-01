package esdata

import "strings"

func removeType(raw string) string {
	split := strings.SplitN(raw, ":", 2)

	if len(split) == 1 {
		return raw
	}

	return split[1]
}

func removeFirstPeriodSeparatedChunk(raw string) string {
	split := strings.SplitN(raw, ".", 2)

	if len(split) == 1 {
		return split[0]
	}

	return split[1]
}
