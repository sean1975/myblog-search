#!/bin/bash

DIR="blog"
OUTPUT="feed.xml"

[ ! -d "$DIR" ] && mkdir -p $DIR

curl -s -o "$DIR/$OUTPUT" http://blog.seanlee.site/feeds/posts/default?max-results=1000

