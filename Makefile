build:
	cd src; go build -o ct *.go

dev:
	cd src; go build -o ct *.go; ./ct --help

clean:
	cd src; rm ct

install:
	cd src; go build -o ct *.go
	cp src/ct /usr/local/bin/ct

