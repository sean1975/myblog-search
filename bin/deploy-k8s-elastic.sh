#!/bin/bash

kubectl create -f kubernetes/myblog-search-env-elastic.yaml

PROVISIONER=`kubectl get sc | grep '\(default\)' | awk '{print $3}' | cut -d/ -f2`

kubectl create -f kubernetes/myblog-search-sc-$PROVISIONER.yaml

kubectl create -f elastic/elastic.yaml

while [[ $(kubectl get pods -l app=myblog-search -l name=elastic -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}' | sort -u) != "True" ]]; do echo "waiting for elastic" && sleep 10; done

kubectl port-forward elastic-0 9200:9200 > /dev/null &
sleep 3

elastic/bin/create_index.sh
[ $? -ne 0 ] && echo "Failed to create index" && jobs && kill %1 && exit 1

elastic/bin/create_search_template.sh
[ $? -ne 0 ] && echo "Failed to create search template" && jobs && kill %1 exit 1

elastic/bin/create_autocomplete_template.sh
[ $? -ne 0 ] && echo "Failed to create autocomplete template" && jobs && kill %1 exit 1

kubectl create -f crawler/crawler-pvc.yaml

kubectl create -f crawler/crawler-batch.yaml

echo "sending a test query to elastic"
kubectl exec elastic-0 -- bash -c 'for i in {1..10}; do sleep 10 && curl -s "http://localhost:9200/_search/template" -H "Content-Type: application/json" -d "{\"id\":\"myblog-search-template\",\"params\":{\"query_string\":\"fish\"}}" | grep -o "\"title\":\[\"[^\"]*\"\]"; if [ $? -eq 0 ]; then echo "successful" && break; fi; echo "retry..."; done'

kubectl delete -f crawler/crawler-batch.yaml

kubectl create -f crawler/crawler-cronjob.yaml

kubectl create -f middleware/middleware.yaml

kubectl create -k nginx

kubectl create -f kubernetes/service.yaml

echo "running end-to-end test"
for i in {1..3}; do sleep 10 && curl -s "http://localhost:80/search/?query=%E9%AD%9A" | grep -o "\"title\":\"[^\"]*\""; if [ $? -eq 0 ]; then echo "successful" && break; fi; echo "retry..."; done

jobs && kill %1
