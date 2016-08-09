package modules

import (
	"reactizer-go/modules/todo"
)

func MountRoutes() {
	todo.Mount()
}
