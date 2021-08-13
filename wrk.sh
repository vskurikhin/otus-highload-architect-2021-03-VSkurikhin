#!/bin/sh
wrk \
  --connections="1" \
  --duration="1m" \
  --script="config.lua" \
  --threads="1" \
  --latency \
  "http://localhost:8080/users/search-by/name/a"
