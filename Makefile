


build: fmt
	go build wemo-control.go
	ln -sf wemo-control off
	ln -sf wemo-control on

fmt:
	go fmt wemo-control.go

install: build
	mkdir -p ~/bin
	cp -pr on off wemo-control ~/bin

clean:
	rm -f on off wemo-control
