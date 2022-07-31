FROM golang:1.18-alpine as builder

#COPY input.json /root
COPY ./* /root/whosbug-CI/
RUN apk update && apk add git
WORKDIR /root/whosbug-CI

ENV GOPROXY="https://proxy.golang.com.cn,direct"

RUN ls -lth
RUN go mod download
RUN go build -o whosbug-CI
RUN ls -lth

FROM golang:1.18-alpine as runner

COPY --from=builder /root/whosbug-CI/whosbug-CI /root/workspace/whosbug-CI
RUN apk update && apk add git

WORKDIR /root/workspace

CMD ["./whosbug-CI"]