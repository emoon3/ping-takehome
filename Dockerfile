FROM alpine:latest
WORKDIR /app
COPY ping /app/
EXPOSE 8080
CMD [ "/app/ping" ]