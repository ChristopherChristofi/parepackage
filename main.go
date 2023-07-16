package main

import (
	"os"
	"log"
	"fmt"
	"flag"
)

func main() {

	var base_req_dir string
	flag.StringVar(&base_req_dir, "base-dir", "data", "path to base directory for requested yaml config files")
	flag.Parse()

	if _, err := os.Stat(base_req_dir); os.IsNotExist(err) {
		log.Fatal(err)
	}

	requested_pkgs := getRequestedPackageNames(base_req_dir)

	for _, pkg := range requested_pkgs {
		fmt.Println(pkg)
	}
}
