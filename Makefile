.PHONY: help
.PHONY: test
.PHONY: run
.PHONY: build
.PHONY: testserver

help:
	@echo "Basic Commands" 
	@echo "test    		=> 'make test'"
	@echo "build   		=> 'make build'"
	@echo "run    		=> 'make run'"
	@echo "testserver 	=> 'make testserver'"
	@echo "godoc"		=> 'make godoc'
	@echo "Run the server and the test server in a seperate terminal instance"

test:
	go test -v ./...

build:
	go build

run: 
	go run main.go

testserver: 
	go run testserver/testserver.go

godoc:
	@echo visit http://localhost:7000/pkg/github.com/koushik-shetty/number-algo-aggregator/
	godoc -http :7000
# function test {
#     go test -v ./...
# }

# case "$1" in
#     "serve")
#     go run main.go
    
#     ;;
#     "prereq")
#     fmPrerequisite
#     ;;
#     "start")
#     fmStart
#     ;;
#     "tests")
#     fmTests
#     ;;
#     "doc")
#     fmStartDoc
#     ;;
#     "info")
#     info
#     ;;
#     *)
#     echo "x Invalid command"
#     echo "Valid commands:"
#     info
#     ;;
# esac
