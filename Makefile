VERSION = "$(shell git rev-parse --short HEAD)"


all:
	go build -ldflags "-X main.VERSION=$(VERSION)" -o fibonize ./fibonize.go
	go build -ldflags "-X main.VERSION=$(VERSION)" -o fizzbuzz ./fizzbuzz.go

run:
	./fibonize | ./fizzbuzz
clean:
	rm -f fibonize fizzbuzz
