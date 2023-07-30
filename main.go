package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func getInstalledPackageNames(pkg_list_path string) (installed_packages []string) {
	list_contents, err := os.Open(pkg_list_path)
	if err != nil {
		log.Fatal(err)
	}
	contentScanner := bufio.NewScanner(list_contents)
	contentScanner.Split(bufio.ScanLines)

	for contentScanner.Scan() {
		installed_packages = append(installed_packages, contentScanner.Text())
	}

	return installed_packages
}

func getRequestedPackageNames(base_dir string) (requested_packages []string) {

	package_yamls, err := os.ReadDir(base_dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range package_yamls {

		contents, err := os.ReadFile(filepath.Join(base_dir, file.Name()))

		if err != nil {
			log.Fatal(err)
		}

		var yaml_config map[string]interface{}
		if err := yaml.Unmarshal(contents, &yaml_config); err != nil {
			log.Fatal(err)
		}

		versions := yaml_config["versions"].(map[string]interface{})

		package_name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		for version := range versions {
			package_name_version := package_name + "@" + version
			requested_packages = append(requested_packages, package_name_version)
		}
	}
	return requested_packages
}

func contains(find_pkg string, pkg_list []string) bool {
	for _, installed_pkg := range pkg_list {
		if installed_pkg == find_pkg {
			return true
		}
	}
	return false
}

func generatePackageStatuses(pkg_list_filename string, base_dir_name string) (installed []string, requested []string) {
	installed = getInstalledPackageNames(pkg_list_filename)
	requested = getRequestedPackageNames(base_dir_name)

	return installed, requested
}

func main() {

	search_command := flag.String("search", "", "showing missing requested packages from installed.")
	file_package_list := flag.String("pkg-list", "data/list.txt", "path to file listing all actually installed packages on system.")
	dir_base_requested := flag.String("base-dir", "data/", "path to base directory for requested yaml config files.")
	flag.Parse()

	if *search_command != "missing" && *search_command != "hidden" {
		log.Fatal("Invalid search option: ", *search_command)
	}

	installed_pkgs, requested_pkgs := generatePackageStatuses(*file_package_list, *dir_base_requested)

	switch *search_command {
	case "missing":
		for _, pkg := range requested_pkgs {
			if !(contains(pkg, installed_pkgs)) {
				fmt.Println(pkg)
			}
		}
	case "hidden":
		for _, pkg := range installed_pkgs {
			if !(contains(pkg, requested_pkgs)) {
				fmt.Println(pkg)
			}
		}
	}
}
