makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

purge:
	rm -f $(makefile_dir)/enumer
	rm -f $(makefile_dir)/cmd/enumer
	rm -f $(makefile_dir)/enum_internal/byte/*.enum.gen.go
	rm -f $(makefile_dir)/enum_internal/string/*.enum.gen.go

build:
	cd $(makefile_dir)/cmd && go build -o $(makefile_dir)/enumer
	cd $(makefile_dir)/cmd && go build -o enumer

generate: 
	go generate ./...

install:
	cd $(makefile_dir)/cmd && go install

test: purge
	@make build
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