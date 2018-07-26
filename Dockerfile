# build stage
FROM golang:1.10-alpine AS build-env
ARG APP_NAME=goapp

RUN apk add --no-cache curl bash git openssh
RUN go get -u github.com/golang/dep/cmd/dep
    
COPY . /go/src/github.com/govinda-attal/hello-world
WORKDIR /go/src/github.com/govinda-attal/hello-world
RUN dep ensure -v && go build -o $APP_NAME

# final stage
FROM alpine:3.7
RUN apk -U add ca-certificates

WORKDIR /app
COPY --from=build-env /go/src/github.com/govinda-attal/hello-world/$APP_NAME /app/

EXPOSE 3333

CMD [ "./hello" ]
#ENTRYPOINT [ "./hello" ]


