# go-bind-plugin-example

This is example plugin that demonstrates how to use [go-bind-plugin](github.com/wendigo/go-bind-plugin). Plugin sources are located in: [plugin/plugin.go](https://github.com/wendigo/go-bind-plugin-example/blob/master/plugin/plugin.go).

Please note that go1.8 (or tip) is required to play with [=buildmode=plugin](https://tip.golang.org/pkg/plugin).

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

## Expected output from [./go-bind-plugin-example](https://github.com/wendigo/go-bind-plugin-example/blob/master/main.go)
```
plug.CalculateSin(100.0) = -0.506366
Hello Gophers!
plug.CurrentYear is: 2016


Wrapper info:
	- Generated on: 2016-11-08 17:58:37.219391162 +0100 CET
	- Command: go-bind-plugin -plugin-path plugin.so -plugin-package ./plugin -output-name PluginAPI -output-path plugin_api.go -output-package  -dereference-vars -rebuild

Plugin info:
	- package: github.com/wendigo/go-bind-plugin-example/plugin
	- sha256 sum: 4f6b6afcc22109fd5532f5a66b24810736a46119633344edecb32297709c04ea
	- size: 2334896 bytes

Exported functions (2):
	- CalculateSin func(float64) (float64)
	- SayHello func(string)

Exported variables (1):
	- CurrentYear int

Plugin imports:
```

As simple as it is :)

See generated [plugin_api.go](https://github.com/wendigo/go-bind-plugin-example/blob/master/plugin_api.go).

Have fun with go plugins!
