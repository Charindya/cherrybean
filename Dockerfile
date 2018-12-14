FROM golang:alpine AS build
COPY main.go main.go
RUN apk update && apk add --no-cache git
RUN go get github.com/nlopes/slack github.com/pkg/errors
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bot main.go

FROM centos
COPY --from=build /bot /go/bin/bot
CMD ["/go/bin/bot"]
