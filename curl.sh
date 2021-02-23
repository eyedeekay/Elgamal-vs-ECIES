#! /usr/bin/env sh

http_proxy=http://127.0.0.1:4444 curl -s -w %{time_total}n -o /dev/null $1
