FROM golang:1.17.6-alpine3.15
RUN apk add build-base

COPY ./backend /opt/build/
WORKDIR /opt/build
RUN go build -o /opt/app/server

CMD sh
FROM alpine:3.15
WORKDIR /opt/app
COPY --from=0 /opt/app/server /opt/app/server
COPY ./backend/index/ /opt/app/index

EXPOSE 8080
CMD /opt/app/server


