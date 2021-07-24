#!/bin/bash

docker-compose -f docker-compose.yml stop
docker rm -f myblog-search-middleware
docker rm -f myblog-search-nginx
docker rm -f myblog-search-vespa
docker volume rm myblog-search_vespa-logs
docker volume rm myblog-search_vespa-var
