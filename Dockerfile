FROM golang:1.15.4

WORKDIR /go/src/go-blog
COPY . /go/src/go-blog
RUN GOPROXY="https://goproxy.io" go build .

EXPOSE 8000
ENTRYPOINT [ "./go-blog" ]