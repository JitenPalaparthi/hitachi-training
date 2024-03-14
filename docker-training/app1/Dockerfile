FROM ubuntu

LABEL Author="Jiten"

LABEL App="Golang App"

WORKDIR /app

# COPY  * /app/

ADD * /app/

RUN chmod 777 service #777 is too open port..

EXPOSE 50090

CMD [ "./service" ]
