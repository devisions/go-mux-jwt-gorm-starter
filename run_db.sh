#!/bin/sh

docker-compose up -d

echo
echo

docker ps --format 'table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}\t{{.Ports}}' -f name=go-starters

echo
echo

sleep 5

docker logs go-starters
