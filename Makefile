VERSION=0.5.0
all:
	docker build --pull=true -t pdu/presents:${VERSION} -f Dockerfile .
