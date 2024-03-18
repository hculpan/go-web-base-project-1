APPNAME=gotemplate
CLINAME=$(APPNAME)-cli

all: build

build:
	templ generate
	go build -o $(APPNAME) cmd/web/*.go
	go build -o $(CLINAME) cmd/cli/*.go

linuxbuild:
	templ generate
	GOOS=linux GOARCH=amd64 go build -o  $(APPNAME).linux cmd/web/*.go
	GOOS=linux GOARCH=amd64 go build -o $(CLINAME).linux cmd/cli/*.go

clean:
	rm -rf  $(APPNAME)
	rm -rf $(CLINAME).*

test:
