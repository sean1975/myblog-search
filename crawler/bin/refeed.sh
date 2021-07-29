#!/bin/bash

BASEDIR=`dirname $0`
BACKUP="backup"

for xmlfile in `find $BACKUP -type f -name "*.xml"`; do
    echo "Re-process $xmlfile"
    $BASEDIR/convert-blog.sh $xmlfile $xmlfile.json
    [ $? -ne 0 ] && echo "Failed to convert $1" && exit 1
    $BASEDIR/feed-blog.sh $xmlfile.json
    [ $? -ne 0 ] && echo "Failed to feed $1" && exit 1
    rm $xmlfile.json
done

echo "Re-fed documents into backend"
