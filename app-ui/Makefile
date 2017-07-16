SRC_FILES += $(shell find src/ -type f)

UI_RESOURCES = node_modules \
               package.json \
               tsconfig.json \
               .angular-cli.json \
               ${SRC_FILES}

all: ui
	@echo "Specify target app to build"

ui: dist/app.asar

dist/app.asar: ${UI_RESOURCES}
	npm run package
	@touch dist/app.asar.stamp

src/shared:
	@echo "Shared resources missing!"

#
# Dev dependencies
#
node_modules:
	@echo "Installing node modules, this can take a while..."
	npm install

#
# Generic cleanup
#
clean:
	$(RM) -r dist

.PHONY: clean
