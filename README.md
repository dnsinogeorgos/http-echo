# http-echo

Echoes http request information.  
Useful for debugging http proxies.  
Incorporates rate limiting.

```
Usage of ./http-echo:
  -a string
        listen address (default ":8080")
  -r int
        number of requests allowed (default 3)
  -w int
        rate window in seconds (default 10)
```

##### install:
```shell
 go get github.com/dnsinogeorgos/http-echo
```
