CMD := cmd/server
TESTS := tests

all:
	cd $(CMD); go build;

test:
	cd $(TESTS); go test;

.PHONY: clean
clean:
	rm -f $(CMD)/server
	find . -name '*~' -delete
	cd $(TESTS); go clean -testcache

