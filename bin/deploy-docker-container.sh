#!/bin/bash

echo "Creating vespa container"
docker run -m 12G --detach --name myblog-search-vespa \
    --hostname myblog-search-vespa \
    --publish 8080:8080 --publish 19112:19112 --publish 19071:19071 \
    sean1975/myblog-search:vespa

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
     sean1975/myblog-search:crawler

echo "Creating docker container for middleware"
docker run -it --detach --name myblog-search-middleware \
     --hostname myblog-search-middleware --publish 8000:80 \
     --env BACKEND_URL="http://host.docker.internal:8080" \
     sean1975/myblog-search:middleware

echo "Creating docker container for frontend"
docker run -it --detach --name myblog-search-nginx \
     --hostname myblog-search-nginx --publish 80:80 \
     sean1975/myblog-search:nginx

echo "Running a test query" && sleep 5
curl -s "http://localhost:80/search/?query=fish"

