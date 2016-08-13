package users

import (
	"database/sql"

	"github.com/kataras/iris"

	"reactizer-go/api/utils"
)

type login struct {
	db *sql.DB
}

type register struct {
	db *sql.DB
}

func (u *register) Serve(c *iris.Context) {
	T := utils.GetT(c)

}