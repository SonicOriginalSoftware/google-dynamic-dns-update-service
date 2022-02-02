.PHONY: clean clean-service clean-all ca all service

FORCE:
.DEFAULT_GOAL := all

OUT_DIR := out

EXE_NAME := google_dynamic_dns_update_service
EXE_PATH := $(OUT_DIR)/$(EXE_NAME)

clean:
	-rm -r $(OUT_DIR)

clean-service:
	-rm $(EXE_PATH)

clean-all: clean-service clean

ca: clean-all

$(EXE_PATH): FORCE
	$(info Building service...)
	go build -o $(EXE_PATH) .

service: $(EXE_PATH)

all: service
