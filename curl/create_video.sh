#!/bin/bash
host='http://localhost:8080'

# get jwt token
token=$(curl -s "$host/jwt" -H 'authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=' | jq -r .token)

# create video
curl -s -X POST "$host/query" \
  -H "authorization: Bearer $token" \
  -H 'content-type: application/json' \
  -d '{"query": "mutation {\n  createVideo(input: {title: \"cidoe 1\", url: \"http://hello.com\", userId: \"1\"}) {\n    id\n    title\n    url\n    author {\n      id\n      name\n    }\n  }\n}"}' \
  | jq
