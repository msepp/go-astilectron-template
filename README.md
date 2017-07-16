# go-astilectron-template

A project template for building desktop apps with Golang, Electron and Angular.

Uses [go-astilectron](https://github.com/asticode/go-astilectron) for Golang/Electron binding.

## Requirements

 * Uses Makefiles for buidling, so you should have `make` installed.
 * `wget` is required for automatically downloading vendored packages
 * `asar` is used for bundling UI resources
  * `npm install -g asar`
 * `angular-cli` for UI scaffolding.
  * See [angular-cli](https://github.com/angular/angular-cli) for installation.
 * Ofcourse Golang.

## Building

When requirements are installed, one should be able to just cd into the directory with main.go and run `make`. This should download all deps, install UI packages and build everything into a single executable.

Building non-host targets happens with `make go-astilectron-template-GOOS-ARCH[.exe]`, where GOOS, ARCH should be replaced with target values.

## Wrinkles
 * Darwin builds are not supported (but maybe are easy to add) since I have no way to verify functionality atm.
 * Uses a `restmpl.go` file for empty resources bundle to make automated builds faster when using IDEs. This file is swapped for the generated resources during build.
 * Does not support development builds for UI atm, but devTools can be enabled/disabled in `Makefile`.
