FROM golang:1.18-alpine as builder

#COPY input.json /root
COPY ./ /root/whosbug-ci/
RUN apk update && apk add git
WORKDIR /root/whosbug-ci

ENV GOPROXY="https://proxy.golang.com.cn,direct"

RUN ls
RUN go build -o whosbug-ci -mod=vendor
RUN ls

FROM golang:1.18-alpine as runner

COPY --from=builder /root/whosbug-ci/whosbug-ci /root/whosbug-ci
RUN apk update && apk add git

WORKDIR /root/workspace

CMD ["/root/whosbug-ci"]