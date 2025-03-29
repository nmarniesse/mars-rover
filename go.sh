#!/bin/bash

docker run -v ./:/app -v ./GOPATH:/go -w /app golang go $@
