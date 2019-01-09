// +build go1.12

package main

import (
	"fmt"
	"runtime/debug"
)

func init() {
	info, available := debug.ReadBuildInfo()
	if !available {
		println("build info not available")
		return
	}
	fmt.Printf("Build info: %+v\n", info)
	if info.Main.Version != "" {
		version = info.Main.Version
	}
}
