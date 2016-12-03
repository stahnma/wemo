
build:
	go build on.go
	ln -sf on off

install: build
	mkdir -p ~/bin
	cp -pr on off ~/bin

