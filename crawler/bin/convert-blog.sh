#!/bin/bash

BASEDIR=`dirname $0`
DIR="blog"
INPUT="$DIR/feed.xml"
OUTPUT="$DIR/feed.json"

[ ! -z "$1" ] && INPUT=$1
[ ! -z "$2" ] && OUTPUT=$2

if [ ! -f "$INPUT" ]; then
  echo "$INPUT does not exist!"
  exit 1
fi

[ -f "$OUTPUT" ] && rm $OUTPUT

$BASEDIR/convert-blog -i $INPUT -o $OUTPUT

