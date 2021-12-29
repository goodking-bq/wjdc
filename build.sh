#!/bin/bash

CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build  -o ynds .
upx -9 ynds
TAG_VERSION=$(git --no-pager log --pretty=oneline --abbrev-commit -1 | awk '{print $1}')
docker build -t harbor.kube.2xi.com/yunwei/ynds:$TAG_VERSION .
docker tag harbor.kube.2xi.com/yunwei/ynds:$TAG_VERSION harbor.kube.2xi.com/yunwei/ynds:latest
echo "push harbor.kube.2xi.com/yunwei/ynds:$TAG_VERSION"
docker push harbor.kube.2xi.com/yunwei/ynds:$TAG_VERSION
echo "push harbor.kube.2xi.com/yunwei/ynds:latest"
docker push harbor.kube.2xi.com/yunwei/ynds:latest
echo "build and push success!"
