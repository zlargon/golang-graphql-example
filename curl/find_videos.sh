#!/bin/bash
host='http://localhost:8080'

# get jwt token
token=$(curl -s "$host/jwt" -H 'authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=' | jq -r .token)

# find all videos
curl -s -X POST "$host/query" \
  -H "authorization: Bearer $token" \
  -H 'content-type: application/json' \
  -d '{"query": "query {\n  videos {\n    id\n    title\n    url\n    author {\n      id\n      name\n    }\n  }\n}\n"}' \
  | jq
