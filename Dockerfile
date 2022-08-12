FROM golang:1.18-alpine as builder

#COPY input.json /root
COPY ./ /root/whosbug-CI/
RUN apk update && apk add git
WORKDIR /root/whosbug-CI

ENV GOPROXY="https://proxy.golang.com.cn,direct"

RUN ls
RUN go build -o whosbug-CI -mod=vendor
RUN ls

FROM golang:1.18-alpine as runner

COPY --from=builder /root/whosbug-CI/whosbug-CI /root/whosbug-CI
RUN apk update && apk add git

WORKDIR /root/workspace

CMD ["/root/whosbug-CI"]