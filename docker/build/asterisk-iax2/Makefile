DOCKER ?= docker
TARGET ?= asterisk-iax2

all:
	@echo "Building ${TARGET}"
	$(DOCKER) build -t $(TARGET) -f Dockerfile .

