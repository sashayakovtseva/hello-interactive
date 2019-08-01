FROM golang:1.12

WORKDIR /hello-interactive
COPY main.go .
RUN go build --ldflags '-linkmode "external" -extldflags "-static"' -o hello main.go

FROM scratch
COPY --from=0 /hello-interactive/hello .
CMD ["./hello"]

