package main

import (
	"github.com/AldieNightStar/mhistea_go/_init"
	"github.com/AldieNightStar/mhistea_go/compiler"
	"github.com/AldieNightStar/mhistea_go/fs"
)

func main() {
	_init.Init()

	tfolder := fs.LocalFs{Path: "D:\\LotImages\\templates"}
	mfolder := fs.LocalFs{Path: "D:\\LotImages\\mods"}
	sfolder := fs.LocalFs{Path: "D:\\LotImages\\story"}

	err := compiler.CompileStory(tfolder, mfolder, sfolder)
	if err != nil {
		println("ERR: " + err.Error())
	}
}
