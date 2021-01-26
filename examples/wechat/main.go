package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var App = hydra.NewApp(
	hydra.WithPlatName("wechat"),
	hydra.WithSystemName("notifier"),
	hydra.WithServerTypes(http.API),
	hydra.WithClusterName("prod"))

func main() {
	bind(App)
	App.Start()
}
