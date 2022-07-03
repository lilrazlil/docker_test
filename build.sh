#!/bin/bash
docker buildx build --platform linux/arm64/v8,linux/amd64 --build-arg ALL="$(cat .git/refs/heads/master)+$(date +%c)" --pull --rm -f "golang-api/Dockerfile" -t go-api:latest "golang-api"
docker buildx build --platform linux/arm64/v8,linux/amd64 --build-arg ALL="$(cat .git/refs/heads/master)+$(date +%c)" --pull --rm -f "golang-db/Dockerfile" -t go-db:latest "golang-db"