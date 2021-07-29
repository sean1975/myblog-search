#!/bin/bash

DIR="blog"
FILE="$DIR/feed.xml"
XSLT="$DIR/feed.xsl"
OUTPUT="$DIR/feed.json"

[ ! -z "$1" ] && FILE=$1
[ ! -z "$2" ] && OUTPUT=$2

if [ ! -f "$FILE" ]; then
  echo "$FILE does not exist!"
  exit 1
fi

[ -f "$OUTPUT" ] && rm $OUTPUT

xsltproc $XSLT $FILE > $OUTPUT

