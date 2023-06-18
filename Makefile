CMD := cmd/server
GUI := gui
TESTS := tests

all:	api/messages.pb.go
	cd $(CMD); go build;

test:	all
	go test github.com/zayaanra/RED/tests/handler

%.pb.go: %.proto
	protoc --go_out=. --go_opt=paths=source_relative $<

.PHONY: clean
clean:
	rm -f $(CMD)/server api/messages.pb.go
	find . -name '*~' -delete
	cd $(TESTS); go clean -testcache

