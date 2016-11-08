# go-bind-plugin-example

This is example plugin that demonstrates how use [go-bind-plugin](github.com/wendigo/go-bind-plugin). Plugin sources are located in: [plugin/](https://github.com/wendigo/go-bind-plugin-example/tree/master/plugin/main.go).

Please note that go1.8 (or tip) is required to play with [plugins](https://tip.golang.org/pkg/plugin).

# How to run/build/play with

```
go version
git clone https://github.com/wendigo/go-bind-plugin-example
cd go-bind-plugin-example
go get github.com/wendigo/go-bind-plugin
go generate
go build
./go-bind-plugin-example
```

As simple as it is :)

See generated [plugin_api.go](https://github.com/wendigo/go-bind-plugin-example/tree/master/plugin/plugin_api.go).

Have fun with go plugins!
