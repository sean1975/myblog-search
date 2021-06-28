#!/bin/bash

DIR="blog"
FILE="feed.xml"
XSLT="feed.xsl"
OUTPUT="feed.json"

if [ ! -f "$DIR/$FILE" ]; then
  echo "$DIR/$FILE does not exist!"
  exit 1
fi

xsltproc $DIR/$XSLT $DIR/$FILE > $DIR/$OUTPUT

