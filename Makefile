

fmt:
	go fmt on.go

build: fmt
	go build on.go
	ln -sf on off

install: build
	mkdir -p ~/bin
	cp -pr on off ~/bin


clean:
	rm -f on off
