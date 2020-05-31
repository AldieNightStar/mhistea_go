package subtext

import (
	"github.com/AldieNightStar/mhistea_go/_common"
	"github.com/AldieNightStar/mhistea_go/sections"
	"testing"
)

const fail = "Test failed!"

func beforeEach() {
	_common.Refs.Sections.SplitToLines = sections.SplitToLines
}

func TestReadSubtext(t *testing.T) {
	beforeEach()

	text := "Line one\nLine two\n\nThis is another subtext\n\nThis is another subtext too\n\n\n\nAnd this is too!"

	subs := ReadSubtext(text)

	if len(subs) != 4 {
		t.Error(fail)
		return
	}
	if subs[0] != "Line one\nLine two" {
		t.Error(fail)
		return
	}
	if subs[1] != "This is another subtext" {
		t.Error(fail)
		return
	}
	if subs[2] != "This is another subtext too" {
		t.Error(fail)
		return
	}
	if subs[3] != "And this is too!" {
		t.Error(fail)
		return
	}
}
