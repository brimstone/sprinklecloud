ARG PACKAGE=github.com/brimstone/sprinklecloud
FROM brimstone/golang:1.16-onbuild as builder

FROM scratch
EXPOSE 8080
ENTRYPOINT ["/sprinklecloud"]
COPY --from=builder /app /sprinklecloud
COPY www /www
