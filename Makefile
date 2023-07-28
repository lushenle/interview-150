ifneq (,)
.error This Makefile requires GNU Make.
endif

GO      = go
GOCOVER = $(GO) tool cover

.PHONY: test 
test:
	$(GO) test -v -cover ./...

.PHONY: update
update:
	$(GO) get -u -v ./...

.PHONY: fmt
fmt:
	$(GO) fmt ./...


.PHONY: cover_test
cover_test:
	$(GO) test -v --coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out


.PHONY: cover_html
cover_html: cover_test
	$(GOCOVER) -html=coverage.out -o coverage.html
	@rm -f coverage.out
	@open coverage.html
