#!/bin/bash

docker-compose -f docker-compose-elastic.yml rm -s -f

while [[ $(curl -s --head http://localhost:9200/myblog | grep "^HTTP.*" | cut -d\  -f2) == "200" ]]; do
    echo "Waiting for elastic cluster"
    sleep 15
done

docker volume rm myblog-search_elastic-data
docker volume rm myblog-search_crawler-backup

