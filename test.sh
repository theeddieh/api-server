#!/bin/sh

curl http://localhost:8080/
curl http://localhost:8080/healthcheck | jq
curl http://localhost:8080/chefs | jq
curl http://localhost:8080/chef/9a784f72-b8a9-4460-a630-0b61234c87bb | jq
curl http://localhost:8080/chef/1bf5f7ba-be25-4873-b602-627e0c2e36c4 | jq