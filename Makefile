.PHONY: clean clean-service clean-all ca all service image openrc

FORCE:
.DEFAULT_GOAL := all

OUT_DIR := out

EXE_NAME := gddns
EXE_PATH := $(OUT_DIR)/$(EXE_NAME)
SERVICE_DIR := $(OUT_DIR)/openrc
SERVICE_PATH := $(SERVICE_DIR)/$(EXE_NAME)
EXE_VERSION := latest

IMAGE_TAG := $(EXE_NAME):$(EXE_VERSION)
IMAGE_PROGRESS := auto
IMAGE_BUILD_TARGET :=
IMAGE_BUILD_TARGET_FLAG :=

ifdef IMAGE_BUILD_TARGET
IMAGE_BUILD_TARGET_FLAG := --target $(IMAGE_BUILD_TARGET)
endif

INSTALL_PREFIX := /usr/local

include service.Makefile

clean:
	-rm -r $(OUT_DIR)

clean-service:
	-rm $(EXE_PATH)

clean-openrc:
	-rm $(SERVICE_PATH)

clean-all: clean-service clean-openrc clean clean-image

ca: clean-all

$(EXE_PATH): FORCE
	$(info Building service...)
	go build -o $(EXE_PATH) .

service: $(EXE_PATH)

$(SERVICE_DIR):
	$(info Creating service directory...)
	mkdir $(SERVICE_DIR)

$(SERVICE_PATH): service.Makefile $(SERVICE_DIR)
	$(info Generating service file...)
	echo "$$SERVICE_CONTENT" > "$(SERVICE_PATH)"
	chmod 0755 $(SERVICE_PATH)

openrc: $(SERVICE_PATH)

install:
	$(info Installing service binary...)
	cp $(EXE_PATH) $(INSTALL_PREFIX)/bin
	$(info Installing service file...)
	cp $(SERVICE_PATH) /etc/init.d/$(EXE_NAME)


all: service openrc
