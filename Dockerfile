FROM golang:1.13.5-stretch

#RUN apt-get update && apt-get -uy upgrade
#RUN apt-get -y install ca-certificates && update-ca-certificates

ENV GO111MODULE=on
RUN mkdir -p /tmp/coredns-delay
COPY . /tmp/coredns-delay
RUN mkdir -p $(go env GOPATH)/src/github.com/tysonvinson/coredns-delay && mv /tmp/coredns-delay $(go env GOPATH)/src/github.com/tysonvinson/coredns-delay
RUN mkdir -p  $(go env GOPATH)/src/github.com/coredns/coredns
RUN git clone https://github.com/coredns/coredns $(go env GOPATH)/src/github.com/coredns/coredns
RUN cd $(go env GOPATH)/src/github.com/coredns/coredns && echo "delay:github.com/tysonvinson/coredns-delay" >> plugin.cfg && make coredns
RUN cp $(go env GOPATH)/src/github.com/coredns/coredns/coredns /coredns
COPY Corefile /

ENTRYPOINT ["/coredns", "-conf", "/Corefile"]
