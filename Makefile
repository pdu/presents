VERSION=0.0.1
all:
	docker build --pull=true -t pdu/presents:${VERSION} -f Dockerfile .
