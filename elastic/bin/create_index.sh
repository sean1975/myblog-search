#!/bin/bash

BACKEND_URL=http://localhost:9200

[ ! -z "$1" ] && BACKEND_URL=$1

curl -X PUT "$BACKEND_URL/myblog?pretty" -H 'Content-Type: application/json' -d'
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0 
  },
  "mappings": {
    "properties": {
      "language": { "type": "keyword" },
      "title": { "type": "text" },
      "url": { "type": "keyword" },
      "body": { "type": "text" },
      "thumbnail": { "type": "keyword" }
    }
  }
}
'

[ $(curl -s --head "$BACKEND_URL/myblog" | grep "^HTTP.*" | cut -d\  -f2) == "200" ] && exit 0
