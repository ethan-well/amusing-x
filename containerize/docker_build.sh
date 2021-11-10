#!/bin/bash

declare -a servers=(
  "amusingplutoserv"
  "amusingcharonserv"
)

version=0.0.22
if [ "$2" != "" ]
then
  version=$2
fi

for s in "${servers[@]}"
do
  docker build --build-arg SERVER_NAME=${s} -t "${s}:${version}" -f "$1" .
done