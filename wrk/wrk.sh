#!/bin/sh

E=1
D=10

for C in 1 10 100 1000
do
  (iostat 10 $((${D}*6)) | tee ./reports/${1}iostat-C$C.txt)&
  T=$(($E % 11))
  echo "T="$T",C="$C
  wrk -t$T -c$C -d${D}m --timeout 30s --latency -s config.lua 'http://localhost:8080' |
  tee ./reports/${1}wrk-C$C.txt
  E=$(($E * 2))
done
