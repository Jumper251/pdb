FROM node:alpine AS builder

WORKDIR /build

COPY ./frontend/package.json ./
RUN npm install
COPY ./frontend .

RUN npm run generate

FROM golang:1.16

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

COPY --from=builder /build/dist /app/frontend/dist

RUN go build -o main .

CMD ["/app/main"]