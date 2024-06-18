makefile_dir	:= $(abspath $(shell pwd))
m				?= "updates"

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

purge:
	rm -f $(makefile_dir)/enumer
	rm -rf $(makefile_dir)/enum_internal/vscode/.vscode/
	find . -name '*.enum.go' -type f -delete

build:
	go mod tidy
	go build $(makefile_dir)/cmd/enumer

generate: 
	go generate ./...

install: build
	go install $(makefile_dir)/cmd/enumer

test: purge build generate
	go test ./...

pull:
	git fetch
	git pull origin master

commit:
	git add . || true
	git commit -m "$(m)" || true

push: commit
	git push origin master

tag:
	git tag -fa $(tag) -m "$(tag)"
	git push -f origin $(tag)

tag-list:
	git fetch --tags
	git tag --list | sort -V

publish: test
	make push m=$(m)
	make tag tag=$(m)
