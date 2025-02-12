#!/bin/bash
docker pull redis:8.0-M03-alpine3.21 
docker run -d -p 6379:6379 redis:8.0-M03-alpine3.21 
