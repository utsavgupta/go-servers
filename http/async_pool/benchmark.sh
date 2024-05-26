#!/usr/bin/sh

GOMAXPROCS=6 go run . 8080 &
p=$!
sleep 2
echo "Benchmarking asynchronous http server with a pool of 5 go routines, where go max procs = 6." | cowsay -f tux | lolcat
wrk -t 4 -c100 -d30s http://localhost:8080
kill -9 $p
fuser -k 8080/tcp

