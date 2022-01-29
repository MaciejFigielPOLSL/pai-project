FROM golang:1.17.6-alpine3.15
RUN apk add build-base

COPY ./backend /opt/build/
WORKDIR /opt/build
RUN go build -o /opt/app/server

FROM node:17-alpine3.14

ENV NODE_OPTIONS=--openssl-legacy-provider

COPY ./frontend /opt/build/
WORKDIR /opt/build
RUN npm run build 

CMD sh
FROM alpine:3.15
WORKDIR /opt/app
COPY --from=0 /opt/app/server /opt/app/server
COPY --from=1 /opt/build/dist /opt/app/index/dist

EXPOSE 8080
CMD /opt/app/server


