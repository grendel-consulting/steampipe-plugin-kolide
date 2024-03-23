STEAMPIPE_INSTALL_DIR ?= ~/.steampipe

# Determine the OS
OS := $(shell uname)

# Is it macOS / Darwin?
ifeq ($(OS),Darwin)
  BUILD_TAGS = netgo
endif

install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/grendel-consulting/kolide@latest/steampipe-plugin-kolide.plugin -tags "$(BUILD_TAGS)" *.go

reconfigure:
	cp config/kolide.spc $(STEAMPIPE_INSTALL_DIR)/config/kolide.spc
