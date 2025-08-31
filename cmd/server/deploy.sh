#!/usr/bin/env sh

git pull
docker compose pull
docker compose build
docker compose up -d
docker compose logs -f

