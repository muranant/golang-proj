FROM golang

ADD *.go /go/src/
RUN go run /go/src/seq.go
RUN cat /tmp/mob*.txt.out|sort

