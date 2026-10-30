.PHONY: build test test-v test-pkg fmt lint clean

build:
	go build ./...

test:
	go test ./...

test-v:
	go test -v ./...

test-pkg:
	@if [ -z "$(PKG)" ]; then \
		echo "Usage: make test-pkg PKG=03-listas/ejercicios/01-ejercicios/..."; \
		exit 1; \
	fi
	go test -v $(PKG)

fmt:
	go fmt ./...

lint:
	@golangci-lint version 2>/dev/null | grep -q 'version 2' || { echo 'Error: se requiere golangci-lint v2.x. Instalalo: https://golangci-lint.run/welcome/install/'; exit 1; }
	golangci-lint run ./...

clean:
	rm -f *.test *.out
