package main

import "tornadoes/app"

func main() {
	app := &app.App{}
	app.Init()
	app.Run(":2814")
}
