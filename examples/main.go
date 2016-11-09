package main

import "github.com/kuloud/gchart"

const start = `version: 1.0
http://localhost:8000`

func main() {
	println(start)
	println(gchart.ListenAndServe(":8000").Error())
}
