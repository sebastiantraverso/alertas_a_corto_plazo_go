# ## BUILD ########################
FROM golang:1.17 as build-env

# # enable Go modules support
ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /go/src/alertas-AC-api

RUN apt-get install --reinstall ca-certificates
RUN mkdir /usr/local/share/ca-certificates/cacert.org
RUN wget -P /usr/local/share/ca-certificates/cacert.org http://www.cacert.org/certs/root.crt http://www.cacert.org/certs/class3.crt
RUN update-ca-certificates
RUN git config --global http.sslCAinfo /etc/ssl/certs/ca-certificates.crt

COPY ./alertas-AC-api /go/src/alertas-AC-api/

RUN go get -d -v /go/src/alertas-AC-api/

RUN go build -o /go/bin/alertas-AC-api
COPY ./alertas-AC-api/config.yaml /go/bin/


# ## DEPLOY #######################
FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/* /
EXPOSE 5000
CMD ["/alertas-AC-api"]

