#!/bin/bash

kubectl delete -f kubernetes/service.yaml

kubectl delete -k nginx

kubectl delete -f middleware/middleware.yaml

kubectl delete -f crawler/crawler-cronjob.yaml

kubectl delete -f crawler/crawler-pvc.yaml

kubectl delete -f elastic/elastic.yaml
kubectl delete pvc elastic-data-elastic-0

PROVISIONER=`kubectl get sc | grep '\(default\)' | awk '{print $3}' | cut -d/ -f2`

kubectl delete -f kubernetes/myblog-search-sc-$PROVISIONER.yaml

kubectl delete -f kubernetes/myblog-search-env-elastic.yaml
