#!/bin/bash

echo "Creating docker container"
docker run -m 12G --detach --name myblog-search-vespa \
    --hostname myblog-search-vespa \
    --publish 8080:8080 --publish 19112:19112 --publish 19071:19071 \
    sean1975/myblog-search:vespa

while [[ $(curl -s --head http://localhost:19071/ApplicationStatus | grep "^HTTP.*" | cut -d\  -f2) != "200" ]]; do
    echo "Waiting for vespa config server"
    sleep 15
done

echo "Creating application package"
(cd vespa/application && zip -r - .) | \
  curl --header Content-Type:application/zip --data-binary @- \
  localhost:19071/application/v2/tenant/default/prepareandactivate
echo ""

while [[ $(curl -s --head http://localhost:8080/ApplicationStatus | grep "^HTTP.*" | cut -d\  -f2) != "200" ]]; do
    echo "Waiting for vespa application"
    sleep 15
done

if [ ! -e bin/vespa-http-client-jar-with-dependencies.jar ]; then
    echo "Downloading feeding tool"
    curl -L -o bin/vespa-http-client-jar-with-dependencies.jar \
        https://search.maven.org/classic/remotecontent?filepath=com/yahoo/vespa/vespa-http-client/7.391.28/vespa-http-client-7.391.28-jar-with-dependencies.jar
fi

echo "Feeding documents"
java -jar bin/vespa-http-client-jar-with-dependencies.jar \
    --verbose --file blog/feed.json --endpoint http://localhost:8080
echo ""

echo "Creating docker container for middleware"
docker run -it --detach --name myblog-search-middleware \
     --hostname myblog-search-middleware --publish 8000:80 \
     --env BACKEND_URL="http://host.docker.internal:8080/search" \
     sean1975/myblog-search:middleware

echo "Creating docker container for frontend"
docker run -it --detach --name myblog-search-nginx \
     --hostname myblog-search-nginx --publish 80:80 \
     sean1975/myblog-search:nginx

echo "Running a test query" && sleep 5
curl -s "http://localhost:80/search/fish"

