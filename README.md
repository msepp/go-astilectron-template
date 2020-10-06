# go-astilectron-template

__NOTE__ this repo is completely redudant and obsolete nowadays and I encourage you not to use this code as anything else than reading material.

A project template for building desktop apps with Golang, Electron and Angular.

Uses [go-astilectron](https://github.com/asticode/go-astilectron) for Golang/Electron binding.

This is meant to be a template, a base for a project that you modify and extend to implement the features you need. As such I won't be covering all different use cases in the code. It might be more insightful to view this as an example than anything else.

## Features
 * Ready recipes for building for Windows/Linux 32/64bit targets
 * Automatically download and install requirements
 * Easily switch Electron version
 * Easily enable/disable dev-tools for builds
 * Builds for different OS/Arch get their specific Electron bundles
 * UI builds are part of the main build
 * Doesn't need extra http server in Go to serve the UI assets
 * Disembed assets either under tmp to leave no files behind or under users home/profile for faster launch times on consecutive launches.

## Requirements

 * You should have [Go](https://golang.org) installed and set up.
 * Uses Makefiles for build automation, so you should have `make` installed.
 * `wget` is required for automatically downloading vendored packages
 * [asar](https://github.com/electron/asar) is used for bundling UI resources
   * `npm install -g asar`
 * [angular-cli](https://github.com/angular/angular-cli) for UI scaffolding.
   * See [angular-cli](https://github.com/angular/angular-cli) for installation.
 * [go-bindata](https://github.com/lestrrat/go-bindata) is used for packing binary files into executable.
   * `go get -u github.com/lestrrat/go-bindata/...`
 * [go-homedir](https://github.com/mitchellh/go-homedir)
   * `go get -u github.com/mitchellh/go-homedir`
 * [go-astilectron](https://github.com/asticode/go-astilectron)
   * `go get -u github.com/asticode/go-astilectron`
 * Windows builds tested using git bash.

## Usage

Make sure you have the requirements installed and run the following commands:

```sh
go get -u github.com/msepp/go-astilectron-template/...
cd $GOPATH/src/github.com/msepp/go-astilectron-template/example
make
```
This will build the sample for your OS/Arch, if supported.

Building non-host targets happens with `make go-astilectron-template-GOOS-ARCH[.exe]`, where GOOS, ARCH should be replaced with target values.

Just copy the example directory and you should now have a working base for a GUI applicaton project.

## Changing electron/astilectron version

Different versions of Electron or Astilectron can be used by specifying the target versions in the main Makefile. When changed, new versions will be downloaded at next build.

## Wrinkles

 * Darwin builds are not supported (but maybe are easy to add) since I have no way to verify functionality atm.
 * Uses a `restmpl.go` file for empty resources bundle to make automated builds faster when using IDEs. This file is swapped for the generated resources during build.
 * Does not support development builds for UI atm, but devTools can be enabled/disabled in `Makefile`.
 * Building on a 32bit host most likely won't work due to limited address space. Use 64bit host and cross-compile.
