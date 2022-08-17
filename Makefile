build:
	mkdir -p bin; cd src; go build -o ct *.go; cp ct ../bin;

dev:
	mkdir -p bin; cd src; go build -o ct *.go; cp ct ../bin; cd ..; ct ./bin/; ./bin/ct --help

clean:
	rm src/ct; rm bin/ct

install:
	cd src; go build -o ct *.go
	cp src/ct /usr/local/bin/ct

