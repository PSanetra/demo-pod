FROM golang:1.13-alpine as build

WORKDIR /src

COPY . /src

RUN go test -v -vet=off ./...

RUN GOOS=linux GARCH=amd64 go build -o demo-pod -ldflags="-s -w" main.go

FROM alpine:3.10

COPY --from=build /src/demo-pod /usr/local/bin

ENTRYPOINT ["demo-pod"]
CMD ["--notes-state-file-path=/mnt/notes/notes.txt", "--watch=demo-secret=/mnt/secrets/demo-secret/secret.txt", "--watch=demo-config-map=/mnt/config-maps/demo-config/config.txt"]
