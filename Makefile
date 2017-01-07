DB_FILE = toy_db.dat
LIB = ./lib/*

test:
	@for pkg in $(wildcard $(LIB)); do go test -v $$pkg; done

run:
	@rm -f $(DB_FILE)
	@go run main.go

repl:
	@rm -f $(DB_FILE)
	@go run repl.go

clean:
	@rm -f $(DB_FILE)

fmt:
	@for pkg in $(wildcard $(LIB)); do go fmt -x $$pkg; done
	@go fmt -x ./*.go
