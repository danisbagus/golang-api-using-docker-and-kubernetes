#!/bin/bash
#

docker run --rm -p 3000:3000  --name app-go danisbagus/go-api

exec "$@"
