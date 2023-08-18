FROM golang:1.21 as builder
ARG NAME
ARG VERSION
ARG BUILD_DATE
ARG BUILD_USER
ARG VCS_REF
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build \
      -o /go/bin/app \
      -ldflags " \
        -X 'github.com/go-deeper/app.name=${NAME}' \
        -X 'github.com/go-deeper/app.version=${VERSION}' \
        -X 'github.com/go-deeper/app.buildDate=${BUILD_DATE}' \
        -X 'github.com/go-deeper/app.buildUser=${BUILD_USER}' \
        -X 'github.com/go-deeper/app.vcsRef=${VCS_REF}' \
      " ./cmd

FROM gcr.io/distroless/base
COPY --from=builder /go/bin/app /app
EXPOSE 8000 8001 8002
CMD ["/app"]
