#!/bin/bash

docker-compose -f docker-compose.yml up -d

while [[ $(curl -s --head http://localhost:19071/ApplicationStatus | grep "^HTTP.*" | cut -d\  -f2) != "200" ]]; do
    echo "Waiting for vespa config server"
    sleep 15
done

echo "Activating application package"
docker exec myblog-search-vespa bash -c '/opt/vespa/bin/vespa-deploy prepare /application && /opt/vespa/bin/vespa-deploy activate'

while [[ $(curl -s --head http://localhost:8080/ApplicationStatus | grep "^HTTP.*" | cut -d\  -f2) != "200" ]]; do
    echo "Waiting for vespa application"
    sleep 15
done

echo "Feeding documents"
docker run -it --rm --name myblog-search-crawler \
     --hostname myblog-search-crawler \
     --env BACKEND_URL="http://host.docker.internal:8080" \
     --volume myblog-search_crawler-backup:/crawler/backup \
     sean1975/myblog-search:crawler

echo "Running a test query" && sleep 5
curl -s "http://localhost:80/search/?query=%E9%AD%9A"

