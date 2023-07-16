package main

import (
	"log"
	"os"
	"bufio"
	"strings"
	"path/filepath"
	"gopkg.in/yaml.v3"
)

func getInstalledPackageNames(pkg_list_path string) (installed_packages []string){
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

func getRequestedPackageNames(base_dir string) (requested_packages []string){

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
			package_name_version :=  package_name + "@" + version
			requested_packages = append(requested_packages, package_name_version)
		}
	}
	return requested_packages
}
