FROM golang:1.16.5

ENV GO111MODULE = auto
ENV GOPROXY=https://goproxy.cn,direct

RUN mkdir /output
WORKDIR /output

ADD . /output

RUN go mod download
RUN go build -o main /output/main.go
EXPOSE 8088
CMD /output/main