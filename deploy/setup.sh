#!/bin/bash
apt-get update
apt-get install golang-go -y
mkdir -p /app/src/github.com/frenata
cd /app/src/github.com/frenata
git clone https://github.com/frenata/api-frenata
cd api-frenata
export GOPATH=/app
go build
touch redirect.json
export PORT=3000
./api-frenata
