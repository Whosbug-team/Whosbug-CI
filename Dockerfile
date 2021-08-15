FROM golang:1.16-alpine

#COPY input.json /root
COPY whosbug_linux_linux /root
RUN apk update && apk add git
WORKDIR /root/workspace
#CMD ls -a
#CMD ["git","clone","git@github.com:Tencent/MMKV.git"]
#ENTRYPOINT ["/root/whosbug_linux_linux"]

#CMD ["/bin/bash"]

#COPY src/whosbugPack/go.mod ./
#COPY src/whosbugPack/go.sum ./
#COPY src/attemp.go ./
#ENV GOPROXY="https://goproxy.io"
#COPY input.json ./
#
#RUN go mod download
#
#COPY src/attemp.go ./
#
#RUN go build -o whosbug-golang
#CMD ["ls"]
##
##CMD ["./whosbug-golang"]