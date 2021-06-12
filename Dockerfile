FROM vespaengine/vespa

ADD image/dict.txt /opt/vespa/conf/dict
ADD image/stopwords.txt /opt/vespa/conf/stopwords

