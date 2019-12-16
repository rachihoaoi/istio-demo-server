FROM alpine:latest

RUN apk add --update ca-certificates
RUN update-ca-certificates
RUN apk add --update tzdata
ENV TZ=Asia/Shanghai
COPY server /
RUN  mkdir /public
RUN chmod +x /server

CMD ["/server"]
