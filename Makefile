.PHONY: clean clean-service clean-all ca all service image

FORCE:
.DEFAULT_GOAL := all

OUT_DIR := out

EXE_NAME := google_dynamic_dns_update_service
EXE_PATH := $(OUT_DIR)/$(EXE_NAME)
SERVICE_PATH := openrc/service
EXE_VERSION := latest

IMAGE_TAG := $(EXE_NAME):$(EXE_VERSION)
IMAGE_PROGRESS := auto
IMAGE_BUILD_TARGET :=
IMAGE_BUILD_TARGET_FLAG :=

ifdef IMAGE_BUILD_TARGET
IMAGE_BUILD_TARGET_FLAG := --target $(IMAGE_BUILD_TARGET)
endif

INSTALL_PREFIX := /usr/local

clean:
	-rm -r $(OUT_DIR)

clean-service:
	-rm $(EXE_PATH)

clean-image-cache:
	-docker image prune -f
	-docker buildx prune -f

clean-image:
	-docker rmi $(IMAGE_TAG)

clean-all: clean-service clean clean-image

ca: clean-all

$(EXE_PATH): FORCE
	$(info Building service...)
	go build -o $(EXE_PATH) .

service: $(EXE_PATH)

image:
	docker buildx build \
		$(IMAGE_BUILD_TARGET_FLAG) \
		--progress $(IMAGE_PROGRESS) \
		-f ./Dockerfile \
		-t $(IMAGE_TAG) \
		.

install:
	$(info Installing service binary...)
	cp $(EXE_PATH) $(INSTALL_PREFIX)/bin
	$(info Installing service file...)
	cp $(SERVICE_PATH) /etc/init.d

all: service
