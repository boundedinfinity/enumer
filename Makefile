makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

clean:
	rm -f $(makefile_dir)/enumer
	rm -f $(makefile_dir)/enum_internal/byte/*.enum.go
	rm -f $(makefile_dir)/enum_internal/string/*.enum.go

install:
	go install

generate:
	go generate ./...

test: clean
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

publish: generate
	make commit m=$(tag)
	make tag tag=$(tag)