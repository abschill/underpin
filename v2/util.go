package main

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
)

type KeyList = []string
type KeyMap2D = []KeyList

type HTMLCVersion struct {
	Major uint8
	Minor uint8
	Patch uint8
}

type HTMLCUtils struct{}

const pkgName = "main"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func LogSection() {
	color.Yellow("%s", "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func LogSubSection() {
	println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func GetType(x interface{}) string {
	return reflect.TypeOf(x).String()
}

// prefixes the packge into the typename ie isType("Foo") checks "main.Foo" bc thats how go lang does it
func IsType(x interface{}, t string) bool {
	return GetType(x) == fmt.Sprintf("%s.%s", pkgName, t)
}

func GetVersion() string {
	runtimeVersion := HTMLCVersion{
		0, 1, 1,
	}
	return fmt.Sprintf("HTMLC Compiler Version: %d.%d.%d\n", runtimeVersion.Major, runtimeVersion.Minor, runtimeVersion.Patch)
}

func GetValidExtensions() []string {
	return []string{
		".htm",
		".html",
		".htmlc",
		".hcl",
	}
}

func GetValidProcessArgs() KeyMap2D {
	return KeyMap2D{
		{
			"-c",
			"--config-file",
		},
		{
			"-l",
			"--loader-file",
		},
		{
			"-d",
			"--debug-file",
		},
	}
}
