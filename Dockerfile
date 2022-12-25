FROM golang:bullseye

RUN git clone https://github.com/sh-serenity/test9.git
WORKDIR test9
RUN make && cp test10 /usr/bin && mkdir /root/.kube
COPY config /root/.kube/config

CMD test10