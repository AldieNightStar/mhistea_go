package section_reader

import "testing"

const TEXT = ":: Home" +
	"\nThis is home!" +
	"\nHaxiDenti programmer" +
	"\n:: Outdoor" +
	"\nThis is outdoor!" +
	"\nMonna Histea"

func TestGetSections(t *testing.T) {
	sections := GetSections(TEXT)

	if len(sections) != 2 {
		t.Fail()
		return
	}
	if sections[0].Name != "Home" || sections[1].Name != "Outdoor" {
		t.Fail()
		return
	}
	if sections[0].LineNumber != 0 || sections[1].LineNumber != 3 {
		t.Fail()
		return
	}
}

func TestSplitToLines(t *testing.T) {
	text := "This\nText\nIs\nSplitted"
	lines := SplitToLines(text)
	if len(lines) != 4 {
		t.Fail()
		return
	}
	if lines[0] != "This" && lines[3] != "Splitted" {
		t.Fail()
		return
	}
}

func TestReadSectionByName(t *testing.T) {
	home := ReadSectionByName(TEXT, "Home")
	outdoor := ReadSectionByName(TEXT, "Outdoor")

	if home != "This is home!\nHaxiDenti programmer\n" {
		t.Fail()
		return
	}
	if outdoor != "This is outdoor!\nMonna Histea\n" {
		t.Fail()
		return
	}
}
