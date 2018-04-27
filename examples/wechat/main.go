package main

import (
	"github.com/micro-plat/hydra/hydra"
)

func main() {
	app := hydra.NewApp(
		hydra.WithPlatName("wechat"),
		hydra.WithSystemName("notifier"),
		hydra.WithServerTypes("api"),
		hydra.WithDebug())
	bind(app)
	app.Start()
}
