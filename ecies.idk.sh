#! /usr/bin/env sh

while true; do
  DATE=$(date)
  mkdir -p $(date +%H) && git add . && git commit -am "Hourly push - $DATE" && git push origin master
  time=$(./curl.sh http://idk.i2p/)
  echo $(date +%H%M): $time | tee -a $(date +%H)/ecies.idk.md
  sleep 4m
done

