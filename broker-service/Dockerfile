FROM alpine:3.17.2

WORKDIR /app

COPY config.yaml .
COPY orbit-app-145b9-firebase-adminsdk-7ycvp-6ab97f8272.json .
COPY brokerApp /

CMD [ "/brokerApp" ]