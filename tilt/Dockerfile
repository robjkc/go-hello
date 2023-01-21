FROM golang:1.13-alpine AS build

RUN apk add --no-cache git
RUN go get github.com/go-delve/delve/cmd/dlv


FROM alpine
EXPOSE 8082 40000
WORKDIR /app
ADD build build
COPY --from=build /go/bin/dlv /dlv
RUN ls -al /
CMD ["bash"]