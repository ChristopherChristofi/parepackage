name_tag ?= parepackage
build_target ?= dev
project_location ?= $(shell pwd)

build:
	docker build --target ${build_target} . --tag ${name_tag}

run:
	docker run -it -v ${project_location}:/project ${name_tag}