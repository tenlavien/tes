package main

import "github.com/tenlavien/spec/tehub"

func main() {
	config := tehub.LoadAppConfig()

	app := tehub.NewApp(config)
	defer app.Stop()

	err := app.Start()
	if err != nil {
		panic(err)
	}
}
