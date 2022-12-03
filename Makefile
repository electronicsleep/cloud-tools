build:
	mkdir -p bin; cd src; go fmt .; go build -o ct *.go
	mv src/ct bin; ./bin/ct --help

test: build
	./bin/ct test test
	./bin/ct echo test -r us-east-1 -e dev
	./bin/ct cs

dev:
	mkdir -p bin; cd src; go build -o ct *.go
	cp src/ct bin; ./bin/ct -h

clean:
	rm src/ct; rm bin/ct

install:
	cd src; go build -o ct *.go
	cp src/ct /usr/local/bin/ct

