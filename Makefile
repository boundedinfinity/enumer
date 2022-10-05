makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

purge:
	rm -f $(makefile_dir)/enumer
	rm -f $(makefile_dir)/enum_internal/byte/*.enum.go
	rm -f $(makefile_dir)/enum_internal/string/*.enum.go

build: generate
	go build

generate: 
	go generate ./...

install: generate
	go install

test: purge
	go build
	go generate ./...
	go test ./...

commit:
	git add . || true
	git commit -m "$(m)" || true
	git push origin master

tag:
	git tag -fa $(tag) -m "$(tag)"
	git push -f origin $(tag)

publish: test
	make commit m=$(m)
	make tag tag=$(m)