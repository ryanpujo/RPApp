FROM alpine:3.17.2

WORKDIR /app

COPY brokerApp /

CMD [ "/brokerApp" ]