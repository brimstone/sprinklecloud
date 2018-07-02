ARG PACKAGE=github.com/brimstone/sprinklecloud
FROM brimstone/golang-musl as builder

FROM scratch
EXPOSE 8080
ENTRYPOINT ["/sprinklecloud"]
COPY --from=builder /app /sprinklecloud
COPY www /www
