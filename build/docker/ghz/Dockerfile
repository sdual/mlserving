FROM golang:1.19-alpine3.18

RUN apk update && apk add git
RUN GOPROXY=direct GOFLAGS="-insecure" go install github.com/bojand/ghz/cmd/ghz@latest
COPY proto/ /proto
COPY small_load_test_1requests.json /
COPY build/ghz/config_1.json /

CMD ["/bin/sh"]
