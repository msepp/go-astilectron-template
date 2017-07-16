# Makefile for building  binaries for different architectures and operating
# systems.
#
# Differences per target:
#  * On Windows, -ldflags "-H windowsgui" is added to hide cmd.exe upon execution
#  * On Windows, .syso file is generated to include icon in built binary
#
# See README.md for details on requirements.
#

# Set to true to have dev tools in build, set to false or disabled for
# prodution builds to omit devtools access
ENABLE_DEV_TOOLS := true

# Set to true to make app unpack assets to temp at lauch. Will cause assets to
# be always unpacked at start and deleted when app closes. If false, assets are
# unpacked under users HOME and vendored assets are re-used between launches.
USE_TEMP := false

# App version is the application version
APP_VERSION = v0.1

# Build string tells where and when the app was built
BUILD_STRING := $(shell git rev-parse --short HEAD)-$(shell git rev-parse --abbrev-ref HEAD)-$(shell hostname)

# Application namespace for temp files etc. generated content
APP_PREFIX := go-astilectron-template

# Application name
APP_NAME := go-astilectron-template ${APP_VERSION}

# App filename is the filename prefix for resulting binaries
APP_FILENAME := go-astilectron-template

# Electron and astilectron versions to use. These versions are passed via env
# and linker to Go, so you should not need to change them anywhere else.
ELECTRON_VERSION := v1.6.11
ASTILECTRON_VERSION := v0.6.0

# Resources dir name
RESOURCES_DIR := resources

# Where to download vendor materials
VENDOR_DIR := ${RESOURCES_DIR}/runtime

# Directory for built UI resources
UI_DIR := ${RESOURCES_DIR}/ui

# Icons directory
ICONS_DIR := ${RESOURCES_DIR}/icons

# Directory for UI sources
UI_SRC_DIR := app-ui

# Path to the final UI bundle
UI_BUNDLE := ${UI_SRC_DIR}/dist/app.asar

# Path to UI index file for setting project name
UI_SRC_INDEX := ${UI_SRC_DIR}/src/index.html

# App .ico path
APP_ICO := ${ICONS_DIR}/app.ico

# Icons used in the final build
ICONS := ${ICONS_DIR}/app.icns \
         ${ICONS_DIR}/app.png \
         ${APP_ICO}

# Paths for the downloaded vendored packages per GOOS and GOARCH.
# Note that astilectron is the same for all platforms.
VENDOR_ASTILECTRON := ${VENDOR_DIR}/astilectron-${ASTILECTRON_VERSION}.zip
VENDOR_LINUX_386 := ${VENDOR_DIR}/electron-${ELECTRON_VERSION}-linux-386.zip
VENDOR_LINUX_AMD64 := ${VENDOR_DIR}/electron-${ELECTRON_VERSION}-linux-amd64.zip
VENDOR_WINDOWS_386 := ${VENDOR_DIR}/electron-${ELECTRON_VERSION}-windows-386.zip
VENDOR_WINDOWS_AMD64 := ${VENDOR_DIR}/electron-${ELECTRON_VERSION}-windows-amd64.zip

# URLs for downloading vendored packages per GOOS and GOARCH.
# Here too we only need one URL for astilectron.
VENDOR_URL_ASTILECTRON := https://github.com/asticode/astilectron/archive/${ASTILECTRON_VERSION}.zip
VENDOR_URL_LINUX_386 := https://github.com/electron/electron/releases/download/${ELECTRON_VERSION}/electron-${ELECTRON_VERSION}-linux-ia32.zip
VENDOR_URL_LINUX_AMD64 := https://github.com/electron/electron/releases/download/${ELECTRON_VERSION}/electron-${ELECTRON_VERSION}-linux-x64.zip
VENDOR_URL_WINDOWS_386 := https://github.com/electron/electron/releases/download/${ELECTRON_VERSION}/electron-${ELECTRON_VERSION}-win32-ia32.zip
VENDOR_URL_WINDOWS_AMD64 := https://github.com/electron/electron/releases/download/${ELECTRON_VERSION}/electron-${ELECTRON_VERSION}-win32-x64.zip

# Bindata builder binary
BINDATA := go-bindata -pkg main -nomemcopy

# Common bindata resources. A list of things to pack into the binary.
# Electron is selected separately based on OS/Arch
# We try to omit all files we don't actually need.
BINDATA_COMMON := ${VENDOR_ASTILECTRON} \
                  ${UI_DIR}/ui.asar \
                  ${ICONS_DIR}/*

# Valid build targets, ie. supported GOOS + GOARCH combinations.
VALID_TARGETS := ${APP_FILENAME}-linux-386 \
                 ${APP_FILENAME}-linux-amd64 \
                 ${APP_FILENAME}-windows-386.exe \
                 ${APP_FILENAME}-windows-amd64.exe

# Filenames for go:generate products per GOOS and GOARCH combo
RESOURCES_LINUX_386 := resources-linux-386.go.res
RESOURCES_LINUX_AMD64 := resources-linux-amd64.go.res
RESOURCES_WINDOWS_386 := resources-windows-386.go.res
RESOURCES_WINDOWS_AMD64 := resources-windows-amd64.go.res

# Common dependencies for go:generate, files that aren't platform dependent but
# are included in the generated .go file.
RESOURCES_DEPS_COMMON := ${UI_DIR}/ui.asar \
                         ${ICONS}

# Common dependencies for the final build product, these are the same for all
# GOOS and GOARCH variants
COMMON_DEPS := Makefile \
               main.go \
               handle_gui.go

# Linker flags for setting version info for config during build.
LDFLAGS := -X 'github.com/msepp/go-astilectron-template/app/bootstrap.electronVersion=${ELECTRON_VERSION}' \
           -X 'github.com/msepp/go-astilectron-template/app/bootstrap.astilectronVersion=${ASTILECTRON_VERSION}' \
           -X 'github.com/msepp/go-astilectron-template/app/bootstrap.appVersion=${APP_VERSION}' \
           -X 'github.com/msepp/go-astilectron-template/app/bootstrap.resourcesDir=${RESOURCES_DIR}' \
           -X 'github.com/msepp/go-astilectron-template/app/bootstrap.devTools=${ENABLE_DEV_TOOLS}' \
           -X 'github.com/msepp/go-astilectron-template/app/bootstrap.useTemp=${USE_TEMP}' \
           -X 'github.com/msepp/go-astilectron-template/app/bootstrap.build=""${BUILD_STRING}""' \
           -X 'github.com/msepp/go-astilectron-template/app/bootstrap.guiName=""${UI_SRC_PROJECT}""' \
           -X 'github.com/msepp/go-astilectron-template/app/bootstrap.name=""${APP_NAME}""' \
           -X 'github.com/msepp/go-astilectron-template/app/bootstrap.prefix=""${APP_PREFIX}""'

DEFAULT_TARGET := ${APP_FILENAME}
ifeq ($(OS),Windows_NT)
  DEFAULT_TARGET := ${DEFAULT_TARGET}-windows
    ifeq ($(PROCESSOR_ARCHITEW6432),AMD64)
      DEFAULT_TARGET := ${DEFAULT_TARGET}-amd64.exe
    else
      ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
        DEFAULT_TARGET := ${DEFAULT_TARGET}-amd64.exe
      else
        ifeq ($(PROCESSOR_ARCHITECTURE),x86)
          DEFAULT_TARGET := ${DEFAULT_TARGET}-386.exe
        else
          DEFAULT_TARGET := unknown
        endif
      endif
   endif
else
  UNAME_S := $(shell uname -s)
  ifeq ($(UNAME_S),Linux)
    DEFAULT_TARGET := ${DEFAULT_TARGET}-linux
    UNAME_P := $(shell uname -p)
    ifeq ($(UNAME_P),x86_64)
      DEFAULT_TARGET := ${DEFAULT_TARGET}-amd64
    else
      ifneq ($(filter %86,$(UNAME_P)),)
        DEFAULT_TARGET := ${DEFAULT_TARGET}-386
      else
        DEFAULT_TARGET := unknown
      endif
    endif
  else
    DEFAULT_TARGET := unknown
  endif
endif

# Template resources file. Just a dummy to get builds working but empty to keep
# build times sensible.
RESTMPL := restmpl.go

# Build instructions with resource file shuffling to use the actual compiled
# resources file instead of placeholder.
BUILD_BIN=@echo "Building $@..." && \
          if [ -f $< ]; then mv $< $(patsubst %.res,%,$<); fi && \
          if [ -f $(RESTMPL) ]; then mv $(RESTMPL) $(RESTMPL).res; fi && \
          GOARCH=$(GOARCH) GOOS=$(GOOS) go build -ldflags "$(LDFLAGS)" -o $@ &&  \
          if [ -f $(patsubst %.res,%,$<) ]; then mv $(patsubst %.res,%,$<) $<; fi && \
          if [ -f $(RESTMPL).res ]; then mv $(RESTMPL).res $(RESTMPL); fi

# All just outputs list of the real targets.
all: ${DEFAULT_TARGET}
	@echo "Built ${DEFAULT_TARGET}"

unknown:
	@echo "Valid targets: ${VALID_TARGETS}"

# Builds binary for 32bit Linux
${APP_FILENAME}-linux-386: GOARCH=386
${APP_FILENAME}-linux-386: GOOS=linux
${APP_FILENAME}-linux-386: ${RESOURCES_LINUX_386} ${COMMON_DEPS}
	@$(RM) rsrc.syso
	$(BUILD_BIN)

# Builds binary for 64bit Linux
${APP_FILENAME}-linux-amd64: GOARCH=amd64
${APP_FILENAME}-linux-amd64: GOOS=linux
${APP_FILENAME}-linux-amd64: ${RESOURCES_LINUX_AMD64} ${COMMON_DEPS}
	@$(RM) rsrc.syso
	$(BUILD_BIN)

# Builds binary for 32bit Windows.
${APP_FILENAME}-windows-386.exe: LDFLAGS += -H windowsgui
${APP_FILENAME}-windows-386.exe: GOARCH=386
${APP_FILENAME}-windows-386.exe: GOOS=windows
${APP_FILENAME}-windows-386.exe: ${RESOURCES_WINDOWS_386} ${COMMON_DEPS} rsrc.syso
	$(BUILD_BIN)

# Builds binary for 64bit Windows
${APP_FILENAME}-windows-amd64.exe: LDFLAGS += -H windowsgui
${APP_FILENAME}-windows-amd64.exe: GOARCH=amd64
${APP_FILENAME}-windows-amd64.exe: GOOS=windows
${APP_FILENAME}-windows-amd64.exe: ${RESOURCES_WINDOWS_AMD64} ${COMMON_DEPS} rsrc.syso
	$(BUILD_BIN)

# Generate .syso file to have an icon for the executable (Windows only).
# This file must be deleted when building for Linux.
rsrc.syso: ${APP_ICO}
	@echo "Generating app icon for Windows target..."
	@rsrc -ico ${APP_ICO}

# Generates resources file for 32bit Linux
${RESOURCES_LINUX_386}: ${VENDOR_LINUX_386} ${BINDATA_COMMON}
	@echo "Creating bindata for Linux 32bit binary. This'll take a while..."
	@$(RM) resources-window-*.go resources-*-amd64.go
	@${BINDATA} -o $@ $^

# Generates resources file for 64bit Linux
${RESOURCES_LINUX_AMD64}: ${VENDOR_LINUX_AMD64} ${BINDATA_COMMON}
	@echo "Creating bindata for Linux 64bit binary. This'll take a while..."
	@$(RM) resources-window-*.go resources-*-386.go
	@${BINDATA} -o $@ $^

# Generates resources file for 32bit Windows
${RESOURCES_WINDOWS_386}: ${VENDOR_WINDOWS_386} ${BINDATA_COMMON}
	@echo "Creating bindata for Windows 32bit binary. This'll take a while..."
	@$(RM) resources-linux-*.go resources-*-amd64.go
	@${BINDATA} -o $@ $^

# Generates resources file for 64bit Windows
${RESOURCES_WINDOWS_AMD64}: ${VENDOR_WINDOWS_AMD64} ${BINDATA_COMMON}
	@echo "Creating bindata for Windows 64bit binary. This'll take a while..."
	@$(RM) resources-linux-*.go resources-*-386.go
	@${BINDATA} -o $@ $^

# Copies bundled UI to the correct place under resources
${UI_DIR}/ui.asar: ${UI_BUNDLE}
	@echo "Copying UI resources..."
	@mkdir -p ${UI_DIR}
	@cp ${UI_BUNDLE} $@

${UI_BUNDLE}: ${UI_BUNDLE}.stamp
	@$(RM) $<

${UI_BUNDLE}.stamp:
	@echo "Updating UI resources..."
	@grep -q '${APP_NAME}' ${UI_SRC_INDEX} || sed -i -e 's/<title>.*<\/title>/<title>${APP_NAME}<\/title>/' ${UI_SRC_INDEX}
	@$(MAKE) -C ${UI_SRC_DIR} ${UI_SRC_PROJECT}

# Downloads astilectron bundle
${VENDOR_ASTILECTRON}:
	@echo "Retrieving astilectron..."
	@mkdir -p ${VENDOR_DIR}
	@wget ${VENDOR_URL_ASTILECTRON} -O $@

# Downloads Electron bundle for 32bit Linux
${VENDOR_LINUX_386}:
	@echo "Retrieving Electron Linux 32bit. This'll take a while..."
	@mkdir -p ${VENDOR_DIR}
	@wget ${VENDOR_URL_LINUX_386} -O $@

# Downloads Electron bundle for 64bit Linux
${VENDOR_LINUX_AMD64}:
	@echo "Retrieving Electron Linux 64bit. This'll take a while..."
	@mkdir -p ${VENDOR_DIR}
	@wget ${VENDOR_URL_LINUX_AMD64} -O $@

# Downloads Electron bundle for 32bit Windows
${VENDOR_WINDOWS_386}:
	@echo "Retrieving Electron Windows 32bit. This'll take a while..."
	@mkdir -p ${VENDOR_DIR}
	@wget ${VENDOR_URL_WINDOWS_386} -O $@

# Downloads Electron bundle for 64bit Windows
${VENDOR_WINDOWS_AMD64}:
	@echo "Retrieving Electron Windows 64bit. This'll take a while..."
	@mkdir -p ${VENDOR_DIR}
	@wget ${VENDOR_URL_WINDOWS_AMD64} -O $@

# Clean removes generated files
clean:
	$(RM) -r ${UI_DIR}
	$(RM) ${VALID_TARGETS}
	$(RM) resources-*.go
	$(RM) resources-*.go.res
	$(RM) rsrc.syso

# Cleans vendored files
clean-vendor:
	$(RM) -r ${VENDOR_DIR}

clean-all: clean-vendor clean

.PHONY:clean clean-vendor
