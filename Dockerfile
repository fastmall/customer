FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates tzdata \
    && rm -rf /var/cache/apk/*
COPY app /app/
ENV ENV prod
RUN mkdir /app/data
VOLUME /app/data
VOLUME /app/conf
WORKDIR /app
CMD ["/app/app"]