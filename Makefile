makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

process:
	go generate ./...

test:
	rm -f $(makefile_dir)/enumer
	go build
	go generate ./...
	go test ./...

commit:
	git add . || true
	git commit -m "$(m)"
	git push origin master

tag:
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)

publish: process
	make commit m=$(tag)
	make tag tag=$(tag)