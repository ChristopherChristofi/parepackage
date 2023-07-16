package main

import (
	"log"
	"os"
	"strings"
	"path/filepath"
	"gopkg.in/yaml.v3"
)

func getRequestedPackageNames(base_dir string) (requested_packages []string){

	requested_packages = []string{}

	package_yamls, err := os.ReadDir(base_dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range package_yamls {

		package_name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		contents, err := os.ReadFile(base_dir + file.Name())
	
		if err != nil {
			log.Fatal(err)
		}
		
		var yaml_config map[string]interface{}
		if err := yaml.Unmarshal(contents, &yaml_config); err != nil {
			log.Fatal(err)
		}

		versions := yaml_config["versions"].(map[string]interface{})
		
		for version := range versions {
			package_name_version :=  package_name + "@" + version
			requested_packages = append(requested_packages, package_name_version)
		}
	}
	return requested_packages
}
