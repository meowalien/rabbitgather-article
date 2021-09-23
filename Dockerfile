FROM golang:1.16.5
WORKDIR "/app"
COPY ../ /app

#ARG jc
#ARG play
ARG IsDebug

ENV Server_Out=bin/server.out

RUN go install
RUN go build -ldflags '-X main.IsDebug='$IsDebug -o $Server_Out .

EXPOSE 2001

ENTRYPOINT ${Server_Out}
