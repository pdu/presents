VERSION=0.4.0
all:
	docker build --pull=true -t pdu/presents:${VERSION} -f Dockerfile .
