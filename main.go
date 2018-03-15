package main

import "io/ioutil"

func main() {

	in, _ := ioutil.ReadFile("./debug.txt")

	parse(string(in))
}
