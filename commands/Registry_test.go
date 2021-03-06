package commands

import "testing"

func TestCommands_UseCommand(t *testing.T) {
	r := NewCommandRegistry()
	AddCommandWithSynonyms(r, []string{"p", "pr", "print"}, "core", "print")

	str := r.UseCommand("print", "\"Hello world!\", true")

	if str == "" {
		t.Fail()
		return
	}
	if str != "mods[\"core\"][\"print\"](\"Hello world!\", true);" {
		t.Fail()
		return
	}
	if r.UseCommand("p", "123") != "mods[\"core\"][\"print\"](123);" {
		t.Fail()
		return
	}
	if r.UseCommand("pr", "\"Hi there\", 1, 2, 3") != "mods[\"core\"][\"print\"](\"Hi there\", 1, 2, 3);" {
		t.Fail()
		return
	}
}

func TestCommands_CommandList(t *testing.T) {
	r := NewCommandRegistry()
	AddCommandWithSynonyms(r, []string{"p", "pr", "print"}, "core", "print")

	list := r.CommandList()

	if len(list) != 3 {
		t.Fail()
		return
	}
	if list[0] != "p" || list[1] != "pr" || list[2] != "print" {
		t.Fail()
		return
	}
}
