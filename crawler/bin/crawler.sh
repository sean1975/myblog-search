#!/bin/bash

BASEDIR=`dirname $0`

$BASEDIR/download-blog.sh
[ $? -ne 0 ] && echo "Failed to download blog" && exit 1

[ ! -f "blog/feed.xml" ] && exit 0

$BASEDIR/convert-blog.sh
[ $? -ne 0 ] && echo "Failed to convert blog" && exit 1

$BASEDIR/feed-blog.sh
[ $? -ne 0 ] && echo "Failed to feed blog" && exit 1

echo "Fed documents into backend"
