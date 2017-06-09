VERSION = 0.2.$(shell git rev-parse --short HEAD)
DATE = $(shell date -u | base64 )
GOBIN=${PWD}/bin/
make: clean
	mkdir -p bin && \
	GOBIN=${GOBIN} go install -v -ldflags='-X main.buildtime="${DATE}" -X main.version="${VERSION}"' ./cmd/...
	@echo build success. now run 'make fizzbuzz' or 'make run' to test
run:
	./bin/fibonize | ./bin/fizzbuzz

fizzbuzz:
	./bin/increment | ./bin/fizzbuzz

clean:
	rm -rvf ./bin/fizzbuzz ./bin/increment ./bin/fibonize
	rmdir ./bin || true

