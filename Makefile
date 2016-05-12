VERSION = "$(shell git rev-parse --short HEAD)"


all:
	mkdir bin || true && go build -ldflags "-X main.VERSION=$(VERSION)" -o bin/fibonize ./fibonize.go && go build -ldflags "-X main.VERSION=$(VERSION)" -o bin/fizzbuzz ./fizzbuzz.go && echo && echo "[fibobuzz]" && echo "Successfully Built" && echo "Now type: make run"

run:
	./fibonize | ./fizzbuzz

