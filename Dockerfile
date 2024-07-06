FROM ubuntu:22.04

RUN apt-get update

WORKDIR /app

COPY . .
RUN chmod 777 host.sh

EXPOSE 8000

CMD ["bash","host.sh"] 

