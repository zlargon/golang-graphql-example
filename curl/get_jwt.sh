#!/bin/bash

curl -s 'http://localhost:8080/jwt' -H 'authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ=' | jq
