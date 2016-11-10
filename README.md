# go-bind-plugin-example

This is example plugin that demonstrates how to use [go-bind-plugin](https://github.com/wendigo/go-bind-plugin). Plugin sources are located in: [plugin/plugin.go](https://github.com/wendigo/go-bind-plugin-example/blob/master/plugin/plugin.go).

Please note that go1.8 (or tip) is required to play with [-buildmode=plugin](https://tip.golang.org/pkg/plugin).

## How to run/build/play with

```
go version
git clone https://github.com/wendigo/go-bind-plugin-example
cd go-bind-plugin-example
go get github.com/wendigo/go-bind-plugin
go generate
go build
./go-bind-plugin-example
```

## Expected output

### go generate
```
go-bind-plugin 12:44:43 Building plugin plugin.so from package ./plugin
go-bind-plugin 12:44:46 Loading and analyzing plugin from: plugin.so
go-bind-plugin 12:44:46 Generating output wrapper: plugin_api.go...
go-bind-plugin 12:44:46 Formatting generated file with gofmt -s -w plugin_api.go
go-bind-plugin 12:44:46 Generated wrapper PluginAPI in file plugin_api.go
```

### [./go-bind-plugin-example](https://github.com/wendigo/go-bind-plugin-example/blob/master/main.go):
```
plug.CalculateSin(100.0) = -0.506366
Hello Gophers!
plug.CurrentYear is: 2016


Struct PluginAPI:
	- Generated on: 2016-11-10 12:41:18.706772009 +0100 CET
	- Command: go-bind-plugin -plugin-path plugin.so -plugin-package ./plugin -output-name PluginAPI -output-path plugin_api.go -output-package main -dereference-vars -rebuild

Plugin info:
	- package: github.com/wendigo/go-bind-plugin-example/plugin
	- sha256 sum: 781e637501b24caa67e2728f47bc2483f7807527dd6cf417cabe359931182002
	- size: 2334896 bytes

Exported functions (2):
	- CalculateSin func(float64) (float64)
	- SayHello func(string)

Exported variables (1):
	- CurrentYear int
```

As simple as it is :)

See generated [plugin_api.go](https://github.com/wendigo/go-bind-plugin-example/blob/master/plugin_api.go).

Have fun with go plugins!
