FROM alpine:3.11 as cert
RUN echo "https://mirrors.aliyun.com/alpine/v3.11/main" > /etc/apk/repositories && \
    apk update && \
    apk --no-cache add ca-certificates

FROM scratch
WORKDIR /
COPY --from=cert /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./gateway"]