FROM golang:1.16-alpine
WORKDIR /app
COPY input.json ./
COPY whosbug_linux ./
RUN apk update && apk add git
CMD ["./whosbug_linux"]

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