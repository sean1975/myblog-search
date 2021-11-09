#!/bin/bash

kubectl create -f kubernetes/myblog-search-env.yaml

PROVISIONER=`kubectl get sc | grep '\(default\)' | awk '{print $3}' | cut -d/ -f2`

kubectl create -f kubernetes/myblog-search-sc-$PROVISIONER.yaml

kubectl create -f vespa/vespa.yaml

while [[ $(kubectl get pods -l app=myblog-search -l name=vespa -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}' | sort -u) != "True" ]]; do echo "waiting for vespa configserver" && sleep 10; done

kubectl cp vespa/application/hosts.xml vespa-0:/application/
kubectl cp vespa/application/services.xml vespa-0:/application/
kubectl cp vespa/application/schemas vespa-0:/application/
kubectl cp vespa/application/components vespa-0:/application/

kubectl exec vespa-0 -- bash -c '/opt/vespa/bin/vespa-deploy prepare /application && /opt/vespa/bin/vespa-deploy activate'

kubectl exec vespa-0 -- bash -c 'while [[ "$(curl -s -o /dev/null -w ''%{http_code}'' http://localhost:8080/ApplicationStatus)" != "200" ]]; do echo "waiting for vespa container" && sleep 10; done'

kubectl create -f crawler/crawler-pvc.yaml

kubectl create -f crawler/crawler-batch.yaml

echo "sending a test query to vespa server"
kubectl exec vespa-0 -- bash -c 'for i in {1..10}; do sleep 10 && curl -s "http://localhost:8080/search/?query=fish" | grep -o "\"totalCount\":[1-9]"; if [ $? -eq 0 ]; then echo "successful" && break; fi; echo "retry..."; done'

kubectl delete -f crawler/crawler-batch.yaml

kubectl create -f crawler/crawler-cronjob.yaml

kubectl create -f middleware/middleware.yaml

kubectl create -k nginx

kubectl create -f kubernetes/service.yaml

echo "running end-to-end test"
for i in {1..3}; do sleep 10 && curl -s "http://localhost:80/search/?query=%E9%AD%9A" | grep -o "\"title\":\"[^\"]*\""; if [ $? -eq 0 ]; then echo "successful" && break; fi; echo "retry..."; done
