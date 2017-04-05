FROM golang:1.8.0
EXPOSE 8020
CMD ["router2"]
COPY . /go/src/github.com/vpcorp/linkerdzipkin/router2
WORKDIR /go/src/github.com/vpcorp/linkerdzipkin/router2
RUN make install