FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN make install
RUN make build
RUN make migrate

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y make

WORKDIR /app


COPY --from=builder /app/shelflove ./
COPY . .
RUN chmod 777 host.sh


EXPOSE 8000

CMD ["sh","host.sh"] 

