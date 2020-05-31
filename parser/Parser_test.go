package parser

import "testing"

func TestParse(t *testing.T) {
	line := "print Hello, World, Me of course"
	args := "Hello, World, Me of course"
	cmd := "print"

	command := Parse(line)
	if command.CommandName != cmd || command.Arguments != args {
		t.Error("Did not parse correctly")
	}
}

func TestParseTemplate(t *testing.T) {
	template := "modules[\"...\"](...);"
	required := "modules[\"core\"](1, 2, 3);"

	actual := ParseTemplate(template, []string{"core", "1, 2, 3"})
	if required != actual {
		t.Error("Template parsing is wrong!")
	}
}

func TestParseTemplateCall(t *testing.T) {
	template := "modules[\"...\"](...);"
	required := "modules[\"core\"](1, 2, 3);"

	actual := ParseTemplateCall(template, "core 1, 2, 3")
	if required != actual {
		t.Error("Template parsing is wrong!")
	}
}
