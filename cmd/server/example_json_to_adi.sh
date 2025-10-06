#!/usr/bin/env sh

curl -v --location 'localhost:8080/' \
--header 'Content-Type: application/x-adif-json' \
--data '{
    "HEADER": {
        "ADIF_VER": "3.1.5"
    },
    "RECORDS": [
        {
            "BAND": "40M",
            "CALL": "K9CTS"
        },
        {
            "BAND": "40M",
            "CALL": "W9PVA"
        }
    ]
}'