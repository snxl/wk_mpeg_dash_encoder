FROM golang:1.18.0-alpine3.15

ENV BENTO4_BIN="/opt/bento4/bin" \
    PATH="$PATH:/opt/bento4/bin"

RUN apk add --update ffmpeg bash curl make

WORKDIR /tmp/bento4

ENV BENTO4_BASE_URL="http://zebulon.bok.net/Bento4/source/" \
    BENTO4_VERSION="1-6-0-639" \
    BENTO4_VERSION_FILE="1-6-0.639" \
    BENTO4_CHECKSUM="5378dbb374343bc274981d6e2ef93bce0851bda1" \
    BENTO4_PATH="/opt/bento4" \
    BENTO4_TYPE="SRC"

RUN apk add --update --upgrade curl python3 unzip bash gcc g++ cmake && \
    curl -O -s ${BENTO4_BASE_URL}/Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}.zip && \
    sha1sum -b Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}.zip && \
    mkdir -p ${BENTO4_PATH} && \
    unzip Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}.zip -d ${BENTO4_PATH} && \
    rm -rf Bento4-${BENTO4_TYPE}-${BENTO4_VERSION_FILE}.zip && \
    apk del unzip && \
    cd ${BENTO4_PATH} && \
    mkdir bin utils && \
    cd ./bin  && cmake -DCMAKE_BUILD_TYPE=Release .. && cmake --build . --config Release && cd .. && \
    cp -R ${BENTO4_PATH}/Source/Python/utils ${BENTO4_PATH}/utils && \
    cp -a ${BENTO4_PATH}/Source/Python/wrappers/. ${BENTO4_PATH}/bin


WORKDIR /go/src

RUN export CGO_CFLAGS="-g -O2 -Wno-return-local-addr"

ENTRYPOINT [ "top" ]