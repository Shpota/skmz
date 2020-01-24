FROM node:12.14 AS JS_BUILD
COPY webapp /webapp
WORKDIR webapp
RUN npm install && npm run build --prod

FROM golang:1.13.6-alpine AS GO_BUILD
COPY server /server
WORKDIR /server
RUN apk add build-base
RUN go build -o /go/bin/server

FROM alpine:3.11
COPY --from=JS_BUILD /webapp/build* ./webapp/
COPY --from=GO_BUILD /go/bin/server ./
CMD ./server
