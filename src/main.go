package main

import "log"

func main() {
	app, err := newSolution()
	if err != nil {
		log.Fatalln(err)
	}
	app.run()
}
