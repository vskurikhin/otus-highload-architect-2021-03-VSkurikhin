#!/bin/sh

E=1
D=10

for C in 1 10 100 1000
do
  T=$(($E % 11))
  echo "T="$T",C="$C
  for host in k8s2 k8s3 k8s4
  do
    ssh root@${host} "(iostat -cktxy $((${D}*6)) 10 >${1}iostat-C$C.txt 2>&1)&">/dev/null 2>&1 &
    ssh root@${host} "(sar -qu -r ALL $((${D}*6)) 10 >${1}sar-C$C.txt)&">/dev/null 2>&1 &
  done
  wrk -t$T -c$C -d${D}m --timeout 30s --latency -s config.lua 'http://localhost:8080' |
  tee ./reports/${1}wrk-C$C.txt
  E=$(($E * 2))
  sleep 60
done
