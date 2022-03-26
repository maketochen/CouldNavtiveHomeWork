FROM golang:1.17-alpine AS Build

WORKDIR /go/src/project/

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"

COPY . /go/src/project/

RUN go build -o /bin/project

FROM alpine:3.10

COPY --from=build /bin/project /bin/project

ENTRYPOINT ["/bin/project"]
