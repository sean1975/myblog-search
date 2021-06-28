#!/bin/bash

DIR="blog"
FILE="feed.json"
if [ -z "$BACKEND_URL" ]; then
    BACKEND_URL="http://localhost:8080"
fi

[[ $(curl -s --head $BACKEND_URL/ApplicationStatus | grep "^HTTP.*" | cut -d\  -f2) != "200" ]] && echo "Backend is not ready" && exit 1

java -jar bin/vespa-http-client-jar-with-dependencies.jar --noretry --verbose --file "$DIR/$FILE" --endpoint "$BACKEND_URL"
