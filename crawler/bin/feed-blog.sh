#!/bin/bash

DIR="blog"
FILE="$DIR/feed.json"
TYPE="vespa"

[ ! -z "$1" ] && FILE=$1
[ ! -z "$2" ] && TYPE=$2

[ ! -f "$FILE" ] && echo "$FILE does not exist!" && exit 1

if [ -z "$BACKEND_URL" ]; then
    BACKEND_URL="http://localhost:8080"
fi

check_vespa_status()
{
    [[ $(curl -s --head $BACKEND_URL/ApplicationStatus | grep "^HTTP.*" | cut -d\  -f2) == "200" ]]
}

check_elastic_status()
{
    [[ $(curl -s --head $BACKEND_URL/myblog | grep "^HTTP.*" | cut -d\  -f2) == "200" ]]
}

check_backend_status()
{
    if [ $TYPE == "vespa" ]; then
        check_vespa_status
    else
        check_elastic_status
    fi
    return $?
}

feed_vespa()
{
    java -jar bin/vespa-http-client-jar-with-dependencies.jar --noretry --verbose --file "$FILE" --endpoint "$BACKEND_URL"
}

feed_elastic()
{
    curl -s -X POST "$BACKEND_URL/_bulk?pretty" -H 'Content-Type: application/x-ndjson' --data-binary "@$FILE" | head -n25
}

feed_backend()
{
    if [ $TYPE == "vespa" ]; then
        feed_vespa
    else
        feed_elastic
    fi
    return $?
}

check_backend_status
if [ $? -ne 0 ]; then
    echo "Backend $TYPE is not ready" && exit 1
fi

feed_backend
if [ $? -ne 0 ]; then
    echo "Failed to feed backend $TYPE" && exit 1
fi

