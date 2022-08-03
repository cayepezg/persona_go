FROM golang AS go_build

WORKDIR /app/
COPY ./ /app/
RUN go mod download && rm -rf ./dist && \
    mkdir -p ./dist/ && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./dist/api_persona ./cmd/main.go


FROM alpine:3.16

WORKDIR /app
COPY --from=go_build  /app/dist/api_persona /app/
COPY --from=go_build /app/view /app/view

EXPOSE 8080
CMD [ "/app/api_persona" ]