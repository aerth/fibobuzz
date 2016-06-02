VERSION = "$(shell git rev-parse --short HEAD)"


make:
	mkdir -p bin && go build -ldflags "-X main.VERSION=$(VERSION)" -o bin/fibonize ./fibonize.go &&\
	go build -ldflags "-X main.VERSION=$(VERSION)" -o bin/fizzbuzz ./fizzbuzz.go && \
	go build -ldflags "-X main.VERSION=$(VERSION)" -o bin/increment ./increment.go
	@echo "[fibobuzz]" && echo "Successfully Built" && echo "Now type: 'make run' or 'make fizzbuzz' to try it out!"

all: build

run:
	echo "Fibonacci | FizzBuzz"
	./bin/fibonize | ./bin/fizzbuzz

fizzbuzz:
	echo "Normal FizzBuzz"
	./bin/increment | ./bin/fizzbuzz

