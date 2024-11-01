all: clean build run

build:
	go build

clean:
	rm -f gameterminal

run:
	./gameterminal
