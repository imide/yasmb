package main

import (
	"github.com/imide/yasmb/pkg/must"
	env "github.com/joho/godotenv"
)

func main() {
	err := env.Load()
	must.Must(err)

}
