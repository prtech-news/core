#!/bin/bash

AWS_HANDLER_PATH=$1
OUTPUT_DIR=$2
argc=$#

if [ -z "$AWS_HANDLER" ] | [ $argc -lt 2 ]
then
  echo "usage: ./aws_build.sh <AWS_HANDLER_PATH> <OUTPUT_DIR>"
else
  echo "Building aws linux executable handler"
  GOOS=linux GOARCH=amd64 go build -o main $AWS_HANDLER_PATH.go
  chmod +x main
  zip main.zip main
  mv main.zip $OUTPUT_DIR/
  mv main $OUTPUT_DIR/
fi