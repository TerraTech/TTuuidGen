TEST := go test -failfast

all: test

.PHONY: fmt
fmt:
	@go fmt

.PHONY: test
test:
	@$(TEST)

.PHONY: testv
testv:
	@$(TEST) -v

.PHONY: testvv
testvv:
	@GOTESTVV=1 $(TEST) -v
