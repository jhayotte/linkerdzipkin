# linkerdzipkin

Welcome to linkerdzipkin documentation. In this repository, I use linkerd as a sidecar container for each service. 

This project will show the advantage of such configuration.


HowTo:
* For each router ```make docker.build```
* ```docker-compose up``` will start our mutualized zipkin and our linkerd in sidecar.
* With ```docker exec linkerd_router1 curl -s 127.1:8080/mymessage``` we send a message through:

----------
| client |
---------
     |
     V
------------------
| linkerd_router1 | ----------------------------
------------------                              |
     |                                          |
     V                                          V
----------                                     ZIPKIN
| router1 |                                     ^
---------                                       |
     |                                          |
     V                                          |
------------------                              |
| linkerd_router2 | ----------------------------
------------------
     | 
     V
----------
| router2 |
---------
     |
     V


and giving the result:

```
$ Router1: root mymessage
$ Router2: received mymessage
```