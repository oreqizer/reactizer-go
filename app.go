package main

import (
	"reactizer-go/server"
	"reactizer-go/modules"
)

func main() {
	modules.MountRoutes()
	server.Start(":8080")
}
