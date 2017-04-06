FROM golang:1.8.0
EXPOSE 8080
CMD ["proxy"]
COPY . /go/src/github.com/vpcorp/linkerdzipkin/proxy
WORKDIR /go/src/github.com/vpcorp/linkerdzipkin/proxy
RUN make install