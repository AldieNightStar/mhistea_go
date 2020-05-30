package packer

import (
	"errors"
	"strings"
)

const ScriptTag = "<!-- script -->"

const ErrorTemplateEmpty = "Error! Template is empty"
const ErrorScriptTagNotFound = "Error! <!-- script --> tag is not found!"
const ErrorWriteFileNotSuccessful = "Error! File writing was not successful!"

//	Packs bundles into <html> file by given template
//	It will replace <!-- script --> to bundles in arguments
func Pack(template, scriptBundle string) (out string, err error) {
	if template == "" {
		return "", errors.New(ErrorTemplateEmpty)
	}
	if !strings.Contains(template, ScriptTag) {
		return "", errors.New(ErrorScriptTagNotFound)
	}
	template = strings.Replace(template, ScriptTag, scriptBundle, 1)
	return template, nil
}

//	Packs bundles into <html> file by given template
//	It will replace <!-- script --> to bundles in arguments
//	Will save result to the file
func PackAndSave(writer FileWriter, fileName, template, scriptBundle string) error {
	packed, err := Pack(template, scriptBundle)
	if err != nil {
		return err
	}
	data := []byte(packed)
	success := writer.WriteFile(fileName, data)
	if success {
		return nil
	} else {
		return errors.New(ErrorWriteFileNotSuccessful)
	}
}
