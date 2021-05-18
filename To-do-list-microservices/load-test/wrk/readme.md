
wrk -t12 -c400 -d30s http://localhost/v1/api/signin -s login.lua
