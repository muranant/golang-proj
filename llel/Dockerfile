FROM golang

ADD *.go /go/src/
RUN mkdir -p /go/src/15-text
ADD 15-text /go/src/15-text/
RUN go run /go/src/parallel.go
RUN cat /tmp/results.txt|sort


