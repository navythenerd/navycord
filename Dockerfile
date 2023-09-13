FROM golang:alpine

WORKDIR /app

COPY . /app/

RUN go mod tidy && go build -o nerdguardian

CMD [ "./nerdguardian" ]