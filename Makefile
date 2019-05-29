PLATFORMS=darwin linux windows

default: deps build
deps:
	@echo "==> Updating build dependencies..."
	go get -u github.com/PuerkitoBio/goquery

build:
	@echo "==> Building for all platforms..."
	$(foreach GOOS, $(PLATFORMS),\
	$(shell GOOS=$(GOOS) GOARCH=amd64 go build -o build/$(GOOS)/goff && \
	tar -czf build/goff-$(GOOS)-amd64.tar.gz build/$(GOOS)/goff))

.PHONY: deps build
