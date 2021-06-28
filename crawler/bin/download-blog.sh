#!/bin/bash

DIR="blog"
OUTPUT="feed.xml"

[ ! -d "$DIR" ] && mkdir -p $DIR

curl -s -o "$DIR/$OUTPUT" http://diaryofsean.blogspot.com/feeds/posts/default?max-results=1000

