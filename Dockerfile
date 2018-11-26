FROM alpine:3.8

# Set Timezone
ENV TIMEZONE Asia/Jakarta
RUN apk update && apk add --no-cache tzdata ca-certificates \
    && cp /usr/share/zoneinfo/`echo $TIMEZONE` /etc/localtime && \
    apk del tzdata

COPY ca-bundle.crt /etc/ssl/certs/ca-certificates.crt
COPY main /
COPY config.yaml /

EXPOSE 3000

CMD ["/main"]