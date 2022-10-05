FORCE:
.DEFAULT_GOAL := all

OUT_DIR := out

EXE_NAME := gddns
EXE_PATH := $(OUT_DIR)/$(EXE_NAME)
SERVICE_DIR := $(OUT_DIR)/openrc
SERVICE_PATH := $(SERVICE_DIR)/$(EXE_NAME)
EXE_VERSION := latest

INSTALL_PREFIX := /usr/local

include service.Makefile

clean-all:
.PHONY: clean-all

clean:
	-rm -r $(OUT_DIR)
clean-all: clean
.PHONY: clean

clean-service:
	-rm $(EXE_PATH)
clean-all: clean-service
.PHONY: clean-service

clean-openrc:
	-rm $(SERVICE_PATH)
clean-all: clean-openrc
.PHONY: clean-openrc

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
.PHONY: openrc

install:
	$(info Installing service binary...)
	cp $(EXE_PATH) $(INSTALL_PREFIX)/bin
	$(info Installing service file...)
	cp $(SERVICE_PATH) /etc/init.d/$(EXE_NAME)
.PHONY: install


all: service openrc
.PHONY: all
