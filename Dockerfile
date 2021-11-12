FROM golang AS builder

COPY . /go/src/mcmanager/

WORKDIR /go/src/mcmanager

RUN go build

FROM openjdk:8u312-jdk-buster

ENV MCMANAGER_PATH=/data

WORKDIR /minecraft

COPY --from=builder /go/src/mcmanager/mcmanager .

ENTRYPOINT [ "./mcmanager" ]
