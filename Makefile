HOSTNAME=github.com
NAMESPACE=justsimplify
NAME=redis-config
BINARY=terraform-provider-${NAME}
VERSION=0.2
OS_ARCH=darwin_amd64

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}