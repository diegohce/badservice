FROM debian:buster-slim AS builder

MAINTAINER Diego Cena <diego.cena@gmail.com>

RUN apt update && apt install -y bzip2 wget

WORKDIR /opt

RUN wget https://github.com/diegohce/badservice/releases/download/v0.1.0/badservice-0.1.0-linux_x64.tar.bz2 && tar -jxvf badservice-0.1.0-linux_x64.tar.bz2 


FROM debian:buster-slim

MAINTAINER Diego Cena <diego.cena@gmail.com>

EXPOSE 6666

WORKDIR /

ENV BADSERVICE_BINDADDR=0.0.0.0:6666

COPY --from=builder /opt/badservice/badservice .

CMD ["./badservice"]

