#!/usr/bin/env bash

PORT=${PORT:-8080}
echo "Calling port ${PORT}"

curl http://localhost:${PORT}/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
curl http://localhost:${PORT}/albums
curl http://localhost:${PORT}/album/2
