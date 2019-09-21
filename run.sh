#!#/bin/bash

# Commad to prompt docker sheel
#docker run -it --rm -v $(pwd):/go/src -w /go/src golang bash

# comman to execute go command
docker run --rm -v $(pwd):/go/src -w /go/src golang go run main.go

  