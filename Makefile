build:
	cd src; go build -o ct *.go

dev:
	mkdir -p bin; cd src; go build -o ct *.go; cp ct ../bin/

clean:
	cd src; rm ct

install:
	cd src; go build -o ct *.go
	cp src/ct /usr/local/bin/ct

