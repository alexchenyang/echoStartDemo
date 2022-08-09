FROM golang:latest

#ENV GO111MODULE = auto
#ENV GOPROXY=https://goproxy.cn,direct

RUN mkdir /output
WORKDIR /output

ADD . /output

#RUN go mod download
RUN export GO111MODULE=auto && export GOPROXY=https://goproxy.cn,direct && go mod download
RUN go build -o main /output/main.go
EXPOSE 8080
CMD /output/main
