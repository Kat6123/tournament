.PHONY: lint

lint:
	gometalinter --enable=misspell --enable=unparam --enable=dupl --enable=gofmt --enable=goimports --disable=gotype --disable=gas --deadline=3m ./...
