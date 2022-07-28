FROM golang as builder
WORKDIR /app
COPY . .
RUN go get && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./onf-test-cli -a -ldflags '-extldflags "-static"' 

FROM gcr.io/distroless/static as runner
COPY --from=builder /app/onf-test-cli /usr/local/bin/onf-test-cli
ENTRYPOINT [ "onf-test-cli" ]