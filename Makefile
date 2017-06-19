VERSION=0.2.0
all:
	docker build --pull=true -t pdu/presents:${VERSION} -f Dockerfile .
