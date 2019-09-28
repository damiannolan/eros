VERSION := 0.0.1

DOCKER_REG = damiannolan
DOCKER_IMAGE = eros
DOCKER_IMAGE_TAG = $(VERSION)
USER = $(shell whoami)

docker-build:
	docker build -t $(DOCKER_REG)/$(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG) .

docker-push:
	docker push $(DOCKER_REG)/$(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG)

docker-build-dev:
	docker build -t $(DOCKER_REG)/$(DOCKER_IMAGE):$(USER) .

docker-push-dev:
	docker push $(DOCKER_REG)/$(DOCKER_IMAGE):$(USER)
	
test:
	go test ./...

test-cover:
	go test ./... -coverprofile=coverage.out && go tool cover -func=coverage.out

.PHONY: \
		docker-build \
		docker-push \
		docker-build-dev \
		docker-push-dev \
		test \
		test-cover