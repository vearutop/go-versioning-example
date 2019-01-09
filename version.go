// +build go1.12

package main

import (
	"fmt"
	"runtime/debug"
)

func init() {
	// delaying init to allow build info setting in another init
	// see https://github.com/golang/go/issues/29628
	// see https://golang.org/src/cmd/go/internal/modload/build.go#L234
	go func() {
		info, available := debug.ReadBuildInfo()
		if !available {
			println("build info not available")
			return
		}
		fmt.Printf("Build info: %+v\n", info)
		if info.Main.Version != "" {
			version = info.Main.Version
		}
	}()
}
