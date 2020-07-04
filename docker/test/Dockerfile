FROM golang:1.13.5-stretch
ENV GO111MODULE=on
WORKDIR apps/iotex-core-rosetta-gateway
COPY . .
RUN git config --global http.proxy 'socks5://192.168.1.6:1080'
RUN git config --global https.proxy 'socks5://192.168.1.6:1080'
RUN export http_proxy=socks5://192.168.1.6:1080
RUN chmod +x tests/test.sh
CMD ./tests/test.sh