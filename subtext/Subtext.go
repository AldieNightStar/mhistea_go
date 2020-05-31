package subtext

import (
	"github.com/AldieNightStar/mhistea_go/_common"
	"strings"
)

func ReadSubtext(text string) []string {
	if text == "" {
		return []string{}
	}
	lines := _common.Refs.Sections.SplitToLines(text)
	builder := strings.Builder{}
	var subTextList []string
	for _, line := range lines {
		if line == "" {
			if builder.Len() > 0 {
				builtString := builder.String()
				for strings.HasSuffix(builtString, "\n") {
					builtString = builtString[0 : len(builtString)-1]
				}
				if len(builtString) != 0 {
					subTextList = append(subTextList, builtString)
				}
			}
			builder.Reset()
			continue
		}
		builder.WriteString(line)
		builder.WriteString("\n")
	}
	if builder.Len() > 0 {
		builtString := builder.String()
		for strings.HasSuffix(builtString, "\n") {
			builtString = builtString[0 : len(builtString)-1]
		}
		subTextList = append(subTextList, builtString)
		builder.Reset()
	}
	return subTextList
}
