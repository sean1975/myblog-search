#!/bin/bash

docker-compose -f deployments/docker-compose.yml stop
docker rm -f myblog-search-middleware
docker rm -f myblog-search-nginx
docker rm -f myblog-search-vespa

