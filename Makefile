VERSION := $(shell git describe --tags)

build:
	go build -o gopt -ldflags "-X main.version=${VERSION}" gopt.go

install: build
	install -d ${DESTDIR}/usr/local/bin/
	install -m 755 ./gopt ${DESTDIR}/usr/local/bin/gopt
