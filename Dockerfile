FROM alpine:3.14 as runtime

USER root

WORKDIR /home/app/

#RUN set -ex \
#    && apk upgrade --no-cache \
#    && apk add \
#    postgresql-client \
#    postgresql-dev \
#    && rm -rf /var/cache/apk/* /tmp/* /var/tmp/*

COPY .env* $ENV_HOME/

COPY ./app .

USER app

CMD ["./app"]
