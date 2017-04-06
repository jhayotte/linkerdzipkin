FROM golang:1.8.0
EXPOSE 8020
CMD ["string"]
COPY . /go/src/github.com/vpcorp/linkerdzipkin/string
WORKDIR /go/src/github.com/vpcorp/linkerdzipkin/string
RUN make install