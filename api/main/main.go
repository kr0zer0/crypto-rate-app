package main

import "api/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
