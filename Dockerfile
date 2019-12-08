FROM node:12-alpine as build-npm

WORKDIR /src

COPY ./ui/package.json ./
COPY ./ui/package-lock.json ./

RUN npm install

COPY ./ui/ ./

RUN npm run build:prod

FROM golang:1.13-alpine as build-go

WORKDIR /src

COPY . /src

RUN go test -v -vet=off ./...

RUN GOOS=linux GARCH=amd64 go build -o demo-pod -ldflags="-s -w" main.go

FROM alpine:3.10

WORKDIR /app

COPY --from=build-go /src/demo-pod ./
COPY --from=build-npm /src/dist/demo-pod-ui ./static

ENTRYPOINT ["/app/demo-pod"]
CMD ["--notes-state-file-path=/mnt/notes/notes.txt", "--watch=demo-secret=/mnt/secrets/demo-secret/secret.txt", "--watch=demo-config-map=/mnt/config-maps/demo-config/config.txt", "--watch=notes=/mnt/notes/notes.txt"]
