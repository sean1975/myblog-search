#!/bin/bash

BACKEND_URL=http://localhost:9200

[ ! -z "$1" ] && BACKEND_URL=$1

[ $(curl -s -X GET "$BACKEND_URL/_scripts/myblog-autocomplete-template" | grep -o '"found":[^,]*' | cut -d: -f2) == "true" ] && exit 0

curl -X PUT "$BACKEND_URL/_scripts/myblog-autocomplete-template?pretty" -H 'Content-Type: application/json' -d'
{
  "script": {
    "lang": "mustache",
    "source": {
      "query": {
        "match_phrase_prefix": {
          "title": {
            "query": "{{query_string}}"
          }
        }
      },
      "fields": [ "title", "url" ],
      "_source": false
    }
  }
}
'

[ $(curl -s -X GET "$BACKEND_URL/_scripts/myblog-autocomplete-template" | grep -o '"found":[^,]*' | cut -d: -f2) == "true" ] && exit 0
