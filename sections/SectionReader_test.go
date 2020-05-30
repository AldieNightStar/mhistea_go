package sections

import "testing"

const TEXT = "This is default text\n" +
	"And this is too!\n" +
	":: Home\n" +
	"This is home!\n" +
	"HaxiDenti programmer\n" +
	":: Outdoor\n" +
	"This is outdoor!\n" +
	"Monna Histea"

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
	if sections[0].LineNumber != 2 || sections[1].LineNumber != 5 {
		t.Fail()
		return
	}
}

func TestSplitToLines(t *testing.T) {
	const text = "This\nText\nIs\nSplitted"
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

func TestReadDefaultSection(t *testing.T) {
	def := ReadDefaultSection(TEXT)
	const expected = "This is default text\nAnd this is too!"

	if len(def) == 0 {
		t.Fail()
		return
	}
	if def != expected {
		t.Fail()
		return
	}
	if len(ReadDefaultSection(":: XXX")) > 0 {
		t.Fail()
		return
	}
}
