package main

import (
	"fmt"
)

func main() {
	requested_pkgs := getRequestedPackageNames()

	for _, pkg := range requested_pkgs {
		fmt.Println(pkg)
	}
}
