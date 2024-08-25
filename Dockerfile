FROM golang:1.17.7-alpine3.15 AS build-env
RUN apk add --update npm
WORKDIR /app
COPY . .
COPY package.json package-lock.json ./
RUN npm install
COPY . ./
RUN npm run build
#RUN apk add --no-cache tree
#RUN tree
ADD . /app
RUN cd /app && go build -o simple-react-frontend

FROM alpine:3.15
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build-env /app/simple-react-frontend /app

EXPOSE 8082
ENTRYPOINT ./simple-react-frontend