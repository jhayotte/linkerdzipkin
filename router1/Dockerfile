FROM golang:1.8.0
EXPOSE 8080
CMD ["router1"]
COPY . /go/src/github.com/vpcorp/linkerdzipkin/router1
WORKDIR /go/src/github.com/vpcorp/linkerdzipkin/router1
RUN make install