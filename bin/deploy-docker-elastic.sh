#!/bin/bash

docker-compose -f docker-compose-elastic.yml up -d

while [[ $(curl -s -X GET "http://localhost:9200/_cluster/health?wait_for_status=yellow&timeout=60s" | grep -o '"timed_out":[^,]*' | cut -d: -f2) != "false" ]]; do
    echo "Waiting for elastic cluster"
    sleep 15
done

BASEDIR=`dirname $0`
ELASTIC=$BASEDIR/../elastic/bin

echo "Creating index"
$ELASTIC/create_index.sh
[ $? -ne 0 ] && echo "Failed to create index" && exit 1

echo "Feeding documents"
docker run -it --rm --name myblog-search-crawler \
     --hostname myblog-search-crawler \
     --env BACKEND_URL="http://host.docker.internal:9200" \
     --env BACKEND_TYPE="elastic" \
     --volume myblog-search_crawler-backup:/crawler/backup \
     sean1975/myblog-search:crawler
[ $? -ne 0 ] && exit 1

echo "Creating search template"
$ELASTIC/create_search_template.sh
[ $? -ne 0 ] && echo "Failed to create search template" && exit 1

echo "Running a test query" && sleep 5
curl -s -X GET "http://localhost:8000/search/?query=fish&pretty"

