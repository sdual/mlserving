FROM golang:1.19-alpine3.18

COPY bin/grpc /
COPY apps/serving/config /config
COPY model.jsonl /

CMD ["/grpc"]
