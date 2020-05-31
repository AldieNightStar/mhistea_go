package packer

import (
	"errors"
	"github.com/AldieNightStar/mhistea_go/_common"
	"strings"
)

const ScriptTag = "<!-- script -->"

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
func PackAndSave(writer _common.FileWriter, fileName, template, scriptBundle string) error {
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
