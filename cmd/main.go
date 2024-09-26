package main

import (
	"github.com/parthvinchhi/db-backup/pkg/routes"
)

func main() {
	r := routes.Routes()

	r.Run(":8080")
}
