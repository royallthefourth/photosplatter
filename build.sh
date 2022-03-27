#!/usr/bin/env sh

set -e

npm ci
npm run build
go build .

echo "Built photosplatter executable."
