# linkerdzipkin

Welcome to linkerdzipkin documentation. In this repository, I use linkerd as a sidecar container for each service. 

This project will show the advantage of such configuration.


In progress:

* With ```docker exec linkerd_router1 curl -s 127.1:8080/mymessage``` we get a message going through:


and giving the result:

```
$ Router1: root mymessage
$ Router2: received mymessage
```