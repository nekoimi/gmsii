FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk update && apk add --no-cache tzdata
ENV TZ Asia/Shanghai
WORKDIR /app
ADD     bin/gmsii_linux_amd64    /app/gmsii
ADD     config.json              /app/config.json
RUN chmod +x /app/gmsii
EXPOSE 8000
ENTRYPOINT ["./gmsii", "server", "--bind=0.0.0.0", "--port=8000", "-c", "config.json"]