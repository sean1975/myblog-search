#!/bin/bash

mvn clean package

cp -f target/myblog-search-*-deploy.jar application/components
