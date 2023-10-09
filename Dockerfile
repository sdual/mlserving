FROM golang:1.19-alpine3.18

COPY bin/grpc /

CMD ["/grpc"]
