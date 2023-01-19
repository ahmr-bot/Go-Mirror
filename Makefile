export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w
build: MirrorsAPI
MirrorsAPI:
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -trimpath -ldflags "$(LDFLAGS)" -o ./out/MirrorsAPI