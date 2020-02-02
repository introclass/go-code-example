.PHONY: install package build
install:
	go install  github.com/jessevdk/go-assets-builder

package:
	go-assets-builder html -o assets.go

build:
	go build
