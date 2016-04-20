VERSION := $(shell git describe --tags)

build:
	go build -o godir -ldflags "-X main.version=${VERSION}" godir.go

install: build
	install -d ${DESTDIR}/usr/local/bin/
	install -m 755 ./godir ${DESTDIR}/usr/local/bin/godir
