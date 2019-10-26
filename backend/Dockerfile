FROM golang:1.13-stretch as builder
WORKDIR /app/src
COPY . .
RUN set -eux \
    && go mod download \
    && CGO_ENABLED=0 go build -o /app/bin .

FROM scratch
COPY --from=builder /app/bin .
ENV PORT=80
EXPOSE $PORT
CMD ["/bin"]
