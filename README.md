# Reverse_Proxy
simple reverse proxy. Round Robin LoadBalancing included

1- Set server.go, reverse.go ADDRESS (ip:port). to use loadbalancing build 2 or more serve

2- Build source code
```
go build <file_name>
```

3- Run
```
./<file_name>
```

4- Test
```
curl ip_addr:port/test
```
Json output like this:
```
{"ip":"RemoteAddress=127.0.0.1:41960, Header=map[Accept:[*/*] Accept-Encoding:[gzip] User-Agent:[curl/7.68.0]]","path":"http://127.0.0.1:4322/test"}
```
