//go:build mage

package main

import (
	"github.com/carolynvs/magex/pkg"
	"github.com/magefile/mage/sh"
)

var Default = Run

const (
	binPath    = "bin"
	binApiPath = "bin/api"
	mainPath   = "./cmd/api"
)

func Build() error {
	return sh.RunV("go", "build", "-o", binApiPath, mainPath)
}

func Run() error {
	return sh.RunV("go", "run", mainPath)
}

func Test() error {
	return sh.RunV("go", "test", "-v", "./...")
}

func Clean() error {
	return sh.RunV("rm", "-rf", binPath)
}

func EnsureMage() error {
	return pkg.EnsureMage("")
}
