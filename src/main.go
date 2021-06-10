package main

import "log"

func main() {
	app := newSolution()
	err := app.run()
	if err != nil {
		log.Fatalln(err)
	}
}
