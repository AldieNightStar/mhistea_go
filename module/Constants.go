package module

const wrappedModScript = "mods[\"@@\"] = (function(){\n" +
	"\tconst mod = {};\n" +
	"\tfunction init(){\n@@\n}\n" +
	"\tinit();\n" +
	"\treturn mod;\n" +
	"})();"
