# linkerdzipkin

Welcome to linkerdzipkin documentation. In this repository, I use linkerd as a sidecar container for each service. 

This project will show the advantage of such configuration.


HowTo:
* ```docker-compose up -d``` will start our mutualized zipkin and our linkerd in sidecar.
* With ```docker exec linkerd_proxy curl -s 127.1:8080/mymessage``` we send a message through:

```
----------
| client |
---------
     |
     V
------------------
| linkerd_proxy | -----------------------------
------------------                             |
     |   ^                                     |
     V   |                                     V
-----------                                 ZIPKIN --> Stores in MYSQL
|   proxy  |                                   ^
-----------                                    |
                                               |
                                               |
------------------                             |
| linkerd_string | ----------------------------
------------------
     |  ^
     V  |
-----------
| string   |
-----------
```

and giving the result:

```
$ proxy: root mymessage
$ string: received mymessage
```

# Zipkin

![alt tag](http://url/to/img.png)

In the picture above you can trace the request of the client and see how long it took for each step.


Current issue:
- Traces between linkerd are not visible inside the same graph..