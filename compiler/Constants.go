package compiler

const sectionFunctionWrapper = "scenes[\"@@\"] = function() {\n" +
	"const subScenes = [];\n" +
	"@@\n" +
	"if (subScenes[0]) subScenes[0]();\n" +
	"}\n"
const subSceneFunctionWrapper = "subScenes[@@] = function() {\n" +
	"const subSceneNumber = @@;\n" +
	"const nextSubScene = () => {\n" +
	"\tif (subScenes[subSceneNumber+1]) subScenes[subSceneNumber+1]()\n" +
	"}\n" +
	"@@\n" +
	"}\n"
const modulesDefinition = "const mods = {};\n"
const scenesDefinition = "const scenes = {};\n"
const startFirstSceneDefinition = "scenes[\"@@\"]()\n"
