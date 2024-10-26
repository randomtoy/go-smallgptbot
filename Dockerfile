FROM --platform=amd64 golang:alpine as builder
ENV VERSION v1.0.0
WORKDIR /app
COPY go.mod go.sum  /app/
RUN go mod download
COPY . .
RUN ls
RUN go build -C cmd/ -buildvcs=false -v -ldflags="-X 'main.Version=$VERSION'"  -o app

FROM --platform=amd64 alpine as prod
WORKDIR /app
COPY --from=builder /app/cmd/app /app
EXPOSE 8080
CMD [ "/app/app" ]