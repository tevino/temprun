all: temprun

.PHONY: clean
clean:
	rm -f temprun temprun.lnx

temprun: app/temprun.go
	CGO_ENABLED=0 go build -o temprun $^

.PHONY: linux
linux: app/temprun.go
	CGO_ENABLED=0 GOOS=linux GO_ARCH=amd64 go build -o temprun.lnx $^

.PHONY: test
test:
	@test/test.sh
	go test ./...

.PHONY: release
release: clean linux
	tar cvf temprun.v${v}.lnx.tar.gz temprun.lnx README.md LICENSE
