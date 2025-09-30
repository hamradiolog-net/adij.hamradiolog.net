#!/usr/bin/env sh

curl -v --location 'localhost:8080/' \
--header 'Content-Type: application/json' \
--data '{
    "header": {
        "ADIF_VER": "3.1.5"
    },
    "records": [
        {
            "band": "40m",
            "call": "K9CTS"
        },
        {
            "band": "40m",
            "call": "W9PVA"
        }
    ]
}'