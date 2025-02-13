FROM chainguard/go as builder
ENV GOOS=linux
ENV CGO_ENABLED=0
ENV GO111MODULE=on
COPY . /app
WORKDIR /app
RUN make testorama

FROM cgr.dev/chainguard/static:latest
COPY --from=builder /app/bin/testorama /testorama
ENTRYPOINT [ "/testorama" ]
