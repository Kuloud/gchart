package main

import "log"
import "github.com/kuloud/gchart"

const start = `version: 1.0
http://localhost:8000`

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	println(start)
	println(gchart.ListenAndServe(":8000").Error())
}
