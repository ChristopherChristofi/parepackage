package main

import (
	"os"
	"log"
	"fmt"
	"flag"
	"golang.org/x/exp/slices"
)

func main() {
	var package_list string
	var base_req_dir string
	flag.StringVar(&package_list, "pkg-list", "data/list.txt", "path to file listing all actually installed packages on system.")
	flag.StringVar(&base_req_dir, "base-dir", "data", "path to base directory for requested yaml config files.")
	flag.Parse()

	if _, err := os.Stat(base_req_dir); os.IsNotExist(err) {
		log.Fatal(err)
	}

	installed_pkgs := getInstalledPackageNames(package_list)
	requested_pkgs := getRequestedPackageNames(base_req_dir)

	for _, pkg := range requested_pkgs {
		if !(slices.Contains(installed_pkgs, pkg)) {
			fmt.Println(pkg)
		}
	}
}
