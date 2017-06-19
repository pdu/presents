#!/bin/bash

VERSION=0.2.0

docker run --name=presents -p 3999:3999 -it --rm -d pdu/presents:$VERSION -http=0.0.0.0:3999 -play=false
