FROM golang AS builder

RUN mkdir /data
RUN mkdir /data/build
COPY ./src /data/src
COPY ./go.mod /data/go.mod

WORKDIR /data

RUN apt update
RUN apt install -y libgl1-mesa-dev xorg-dev

RUN go mod tidy && go mod download
RUN go build -o build/auth_server src/auth/cmd/*.go

FROM golang

COPY --from=builder /data/build/auth_server /auth_server

ENV RUN_IN_DOCKER=1

RUN apt update
RUN apt install -y libgl1-mesa-dev xorg-dev
RUN mkdir logs

ENTRYPOINT ["/auth_server"]