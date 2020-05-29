package parser

import "strings"

type Command struct {
	CommandName string
	Arguments   string
}

const TEMPLATE_STRING = "..."

//	Parse string line command into Command { CommandName, Arguments }
//		CommandName string - is a name of the Command
//		Arguments string - is actual arguments to the Command
func Parse(line string) (c Command) {
	c = Command{}
	if line == "" {
		return
	}
	if strings.Contains(line, " ") {
		arr := strings.SplitN(line, " ", 2)
		c.CommandName = arr[0]
		c.Arguments = strings.Trim(arr[1], " \t")
	} else {
		c.CommandName = line
	}
	return c
}

//	Parse string command line into templated string.
//	Example:
//		text := ParseTemplateCall("...(...)", "print 'Hello!'")
//		println(text) // Output will be: print('Hello!')
func ParseTemplateCall(template, line string) (out string) {
	if line == "" {
		return ""
	}
	cmd := Parse(line)
	out = ParseTemplate(template, []string{cmd.CommandName, cmd.Arguments})
	return out
}

//	Parse []string text into templated string.
//	Example:
//		text := ParseTemplate("My name is ... and i am ... years old", []string{"Haxi", "32"})
//		println(text) // Output will be: My name is Haxi and i am 32 years old
func ParseTemplate(template string, text []string) (out string) {
	out = template
	for i := 0; i < len(text); i++ {
		out = strings.Replace(out, TEMPLATE_STRING, text[i], 1)
	}
	return out
}
