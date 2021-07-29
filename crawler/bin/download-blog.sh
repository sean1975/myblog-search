#!/bin/bash

DIR="blog"
BACKUP="backup"
OUTPUT="$DIR/feed.xml"

[ ! -d "$DIR" ] && mkdir -p $DIR
[ ! -d "$BACKUP" ] && mkdir -p $BACKUP

PREVIOUS=`find $BACKUP -type f -name "*.xml" | sort | tail -n1`

QUERY_PARAMETERS="max-results=1000"

if [ ! -z "$PREVIOUS" ]; then
    echo "Blog was downloaded at $PREVIOUS"
    QUERY_PARAMETERS="$QUERY_PARAMETER&orderby=updated&updated-min=$PREVIOUS"
fi

curl -s -o "$OUTPUT" "http://blog.seanlee.site/feeds/posts/default?$QUERY_PARAMETERS"

COUNT=`xmllint --xpath "count(//*[name() = 'feed']/*[name() = 'entry'])" $OUTPUT`

[ $COUNT -eq 0 ] && echo "Blog is up-to-dated" && rm $OUTPUT && exit 0

TIMESTAMP=`xmllint --xpath "//*[name() = 'feed']/*[name() = 'entry']/*[name() = 'updated']/text()" $OUTPUT | sed "s/+[0-9][0-9]:[0-9][0-9]/\n/g" | sort | tail -n1`

[ -z "$TIMESTAMP" ] && echo "Failed to parse timestamp in $OUTPUT" && exit 1

echo "Blog was updated at $TIMESTAMP"

cp $OUTPUT $BACKUP/$TIMESTAMP.xml
