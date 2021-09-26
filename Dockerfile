FROM golang:1.16.5


WORKDIR "/app"
COPY ./ ./

ENV DEBUG="false"

RUN cd /app/sec/ && go install && go build -o /app/bin/server.out .
# RUN go install
# RUN go build /app/sec/ -o /app/bin/server.out .

EXPOSE 2001

ENTRYPOINT cd /app/sec/ && /app/bin/server.out -debug=$DEBUG
