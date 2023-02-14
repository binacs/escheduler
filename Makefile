BUILD_TAGS?=
BUILD_FLAGS = -ldflags "-X github.com/binacs/escheduler/version.GitCommit=`git rev-parse HEAD`"

default: clean build

clean:
	rm -rf bin

build:
	go build $(BUILD_FLAGS) -tags '$(BUILD_TAGS)' -o bin/escheduler ./cmd

mock:
	cd core && go generate; cd -
	cd framework && go generate; cd -

docker:
	docker build -t binacs/escheduler:latest . 

test:
	go test ./... -cover

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: mock test