#! /usr/bin/env sh

while true; do
  DATE=$(date)
  mkdir $(date +%H) && git add . && git commit -am "Hourly push - $DATE" && git push origin main
  time=$(./curl.sh http://mwayujrtqsxzjtvjr53f6x7ioxhecujsdqs6acjhzjtyhoo3eaaa.b32.i2p/)
  echo $(date +%H%M): $time | tee -a $(date +%H)/elg.idk.md
  sleep 4m
done
