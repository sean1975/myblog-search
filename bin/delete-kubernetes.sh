#!/bin/bash

SRC_DIRECTORY=kubernetes

kubectl delete -f ${SRC_DIRECTORY}/service.yaml

kubectl delete -f ${SRC_DIRECTORY}/nginx.yaml

kubectl delete -f ${SRC_DIRECTORY}/middleware.yaml

kubectl delete -f ${SRC_DIRECTORY}/crawler.yaml

kubectl delete -f ${SRC_DIRECTORY}/vespa.yaml
kubectl delete pvc vespa-application-vespa-0
kubectl delete pvc vespa-conf-jieba-vespa-0
kubectl delete pvc vespa-logs-vespa-0
kubectl delete pvc vespa-var-vespa-0

kubectl delete -f ${SRC_DIRECTORY}/configmap/myblog-search-env.yaml
