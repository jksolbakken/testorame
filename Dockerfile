FROM golang:1.19.4-alpine as builder
RUN apk add --no-cache make
ENV GOOS=linux
ENV CGO_ENABLED=0
ENV GO111MODULE=on
COPY . /app
WORKDIR /app
RUN make testorama

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /app/bin/testorama /testorama
ENTRYPOINT [ "/testorama" ]
