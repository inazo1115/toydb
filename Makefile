DB_FILE = toy_db.dat
LIB = ./lib/*

help:
	@echo This is the Makefile of toy_db.
	@echo Available commands are below.
	@cat Makefile \
		| grep -o -e '^[a-z]\+:' \
		| grep -v grep \
		| sed 's/://g' \
		| while read cmd; do echo "- $$cmd"; done

run:
	@rm -f $(DB_FILE)
	@go run main.go

clean:
	@rm -f $(DB_FILE)

test:
	@for pkg in $(wildcard $(LIB)); do go test -v $$pkg; done

fmt:
	@for pkg in $(wildcard $(LIB)); do go fmt -x $$pkg; done
	@go fmt -x ./*.go
