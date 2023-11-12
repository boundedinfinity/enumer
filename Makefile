makefile_dir	:= $(abspath $(shell pwd))
m				?= "updates"

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

purge:
	rm -f $(makefile_dir)/enumer
	rm -f $(makefile_dir)/enum_internal/string/*.enum.go
	rm -rf $(makefile_dir)/enum_internal/vscode/.vscode/

build:
	go build $(makefile_dir)/cmd/enumer

generate: 
	go generate ./...

install: build
	go install $(makefile_dir)/cmd/enumer

test: purge build generate
	go test ./...

push:
	git add . || true
	git commit -m "$(m)" || true
	git push origin master

tag:
	git tag -fa $(tag) -m "$(tag)"
	git push -f origin $(tag)

publish: test
	make push m=$(m)
	make tag tag=$(m)
