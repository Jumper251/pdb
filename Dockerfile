FROM node:alpine AS builder

WORKDIR /build

ARG BASE_URL
ENV BASE_URL=${BASE_URL}

RUN echo "value for BASE_URL: [${BASE_URL}]"

COPY ./frontend/package.json ./
RUN npm install
COPY ./frontend .

RUN npm run generate

FROM golang:1.16

WORKDIR /app

ENV GIN_MODE=release

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

COPY --from=builder /build/dist /app/frontend/dist

RUN go build -o main .

CMD ["/app/main"]