#!/bin/bash

docker run --name=presents -p 3999:3999 -it --rm -d pdu/presents:0.0.1 -http=0.0.0.0:3999 -play=false
