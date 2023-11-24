FROM golang:1.21-alpine3.18

WORKDIR /app
COPY . .

ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_USERNAME=${DB_USERNAME}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_DATABASE=${DB_DATABASE}
ENV PORT=${PORT}

RUN go mod tidy
RUN go mod download
RUN go build -o main .

CMD [ "./main" ]
EXPOSE ${PORT}
