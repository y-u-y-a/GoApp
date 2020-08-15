FROM golang:1.13-alpine
# appディレクトリ生成
WORKDIR /code

# 必要なライブラリインストール
RUN apk add --no-cache git gcc libc-dev

RUN go get github.com/oxequa/realize
RUN go get github.com/labstack/echo

CMD ["realize", "start"]
