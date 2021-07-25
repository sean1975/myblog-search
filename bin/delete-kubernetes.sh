#!/bin/bash

kubectl delete -f kubernetes/service.yaml

kubectl delete -k nginx

kubectl delete -f middleware/middleware.yaml

kubectl delete -f crawler/crawler.yaml

kubectl delete -f vespa/vespa.yaml
kubectl delete pvc vespa-application-vespa-0
kubectl delete pvc vespa-conf-jieba-vespa-0
kubectl delete pvc vespa-logs-vespa-0
kubectl delete pvc vespa-var-vespa-0

kubectl delete -f kubernetes/myblog-search-env.yaml
