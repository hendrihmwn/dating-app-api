FROM golang:1.21.3
ENV HOME="/opt/app-root"
RUN set -xe \
 && apt-get update -qy \
 && apt-get install -qy software-properties-common gettext-base git \
 && mkdir -p ${HOME}
WORKDIR ${HOME}
COPY . .
RUN make prepare
RUN make install
RUN set -xe \
 && chown -R 1001 . \
 && chgrp -R 0 . \
 && chmod -R g=u .
USER 1001
EXPOSE 8000
CMD ["sh", "-c",  "dating-app"]