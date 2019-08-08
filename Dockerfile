FROM sifar/isu7:0.0.1

RUN apt-get install wget unzip

ENV GOROOT /home/isucon/local/go
ENV GOPATH $HOME/../isubata/webapp/go
ENV ISUBATA_DB_PASSWORD isucon
RUN export PATH=/home/isucon/local/go/bin:$PATH

EXPOSE 80
