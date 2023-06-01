# build stage
FROM golang:1.20-alpine3.17 AS builder
LABEL maintainer="prasanthpn <prasanthpn68@gmail.com"
RUN apk update &&  apk add --no-cache git 
WORKDIR /project
COPY . .
RUN apk add --no-cache make 
RUN make deps
RUN go mod vendor
RUN make build
#run stage 
FROM alpine:3.17
WORKDIR /project
COPY  .env .
COPY templates ./templates
COPY --from=builder /project/build/bin/api .
EXPOSE 8080
CMD [ "/project/api" ]
# CMD [ "make","run" ]