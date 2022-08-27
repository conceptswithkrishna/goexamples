ARG ALPINE_VERSION=3.13
FROM alpine:${ALPINE_VERSION}

RUN apk --update --no-cache add ca-certificates go build-base curl automake autoconf git libtool zlib-dev

RUN addgroup -g 1000 protoc && adduser -u 1000 -S -G protoc protoc

RUN mkdir -p /tmp/protobuf

WORKDIR /tmp

ARG PROTOBUF_VERSION=3.21.5
RUN curl -L https://github.com/google/protobuf/archive/v${PROTOBUF_VERSION}.tar.gz | tar xvz --strip-components=1 -C /tmp/protobuf

RUN cd protobuf && autoreconf -f -i -Wall,no-obsolete && \
    ./configure --prefix=/usr --enable-static=no && \
    make -j2 && make install

RUN rm -rf /tmp/protobuf

ENV GOPATH=/go

RUN mkdir -p ${GOPATH} ${GOPATH}/src/github.com/golang/protobuf

# Protoc for Golang
ARG PROTOC_GEN_GO_VERSION=1.4.3
RUN curl -sSL https://api.github.com/repos/golang/protobuf/tarball/v${PROTOC_GEN_GO_VERSION} | tar xz --strip 1 -C ${GOPATH}/src/github.com/golang/protobuf
   
WORKDIR  ${GOPATH}/src/github.com/golang/protobuf

RUN go build -ldflags '-w -s' -o /golang-protobuf-out/protoc-gen-go ./protoc-gen-go

RUN install -Ds /golang-protobuf-out/protoc-gen-go /usr/bin/protoc-gen-go

RUN chmod a+x /usr/bin/protoc

RUN mkdir -p /proto /go-out

# Protoc validation
ARG PROTOC_GEN_VALIDATE_VERSION=0.6.0
RUN mkdir -p ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate && \
    curl -sSL https://api.github.com/repos/envoyproxy/protoc-gen-validate/tarball/v${PROTOC_GEN_VALIDATE_VERSION} | tar xz --strip 1 -C ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate && \
    cd ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate && make build
WORKDIR ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate
RUN go mod download
RUN go build -ldflags '-w -s' -o /protoc-gen-validate-out/protoc-gen-validate .
RUN install -Ds /protoc-gen-validate-out/protoc-gen-validate /usr/bin/protoc-gen-validate
RUN install -D ./validate/validate.proto /out/usr/include/github.com/envoyproxy/protoc-gen-validate/validate/validate.proto

RUN apk del ca-certificates go curl automake autoconf git libtool zlib-dev

ENTRYPOINT [ "/usr/bin/protoc", "-I=/proto", "-I=/go/src", "-I=/go/src/github.com/envoyproxy/protoc-gen-validate", "--go_out=/go-out", "--validate_out=lang=go:/go-out"]

USER protoc:protoc

