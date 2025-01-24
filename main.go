package main

import "loompkg/cli"

func main() {
	if err := cli.Command().Execute(); err != nil {
		panic(err)
	}
}
