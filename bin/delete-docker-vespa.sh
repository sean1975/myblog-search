#!/bin/bash

docker-compose -f docker-compose-vespa.yml stop
docker rm -f myblog-search-middleware
docker rm -f myblog-search-nginx
docker rm -f myblog-search-vespa
docker volume rm myblog-search_vespa-logs
docker volume rm myblog-search_vespa-var
docker volume rm myblog-search_crawler-backup
