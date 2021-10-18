#!/bin/bash

BACKEND_URL=http://localhost:9200

[ ! -z "$1" ] && BACKEND_URL=$1

curl -X PUT "$BACKEND_URL/_scripts/myblog-search-template?pretty" -H 'Content-Type: application/json' -d'
{
  "script": {
    "lang": "mustache",
    "source": {
      "query": {
        "multi_match": {
          "query": "{{query_string}}",
          "fields": [ "title", "body" ]
        }
      },
      "fields": [ "title", "url", "thumbnail" ],
      "highlight": {
        "fields": {
          "body": {}
        }
      },
      "_source": false
    }
  }
}
'

[ $(curl -s -X GET "$BACKEND_URL/_scripts/myblog-search-template" | grep -o '"found":[^,]*' | cut -d: -f2) == "true" ] && exit 0
