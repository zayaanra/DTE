CMD := cmd/server
FRONTEND := ../frontend
TESTS := tests

all:	api/messages.pb.go
	cd $(CMD); go build;

test:	all
	go test github.com/zayaanra/RED/backend/tests/handler

%.pb.go: %.proto
	protoc --go_out=. --go_opt=paths=source_relative $<
	cd $(FRONTEND); protoc -I=$(FRONTEND) --python_out=$(FRONTEND) $(FRONTEND)/messages.proto

.PHONY: clean
clean:
	rm -f $(CMD)/server api/messages.pb.go $(FRONTEND)/messages_pb2.py
	find . -name '*~' -delete
	cd $(TESTS); go clean -testcache

