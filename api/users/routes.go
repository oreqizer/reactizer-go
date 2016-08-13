package users

import (
	"database/sql"
	"log"

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
	user := &User{}
	err := c.ReadJSON(user)
	if err != nil {
		log.Print(err)
		return
	}

	err = checkUsername(user.Username, u.db)
	if err != nil {
		c.Error(T(err.Error()), 409)
		return
	}
	err = checkEmail(user.Email, u.db)
	if err != nil {
		c.Error(T(err.Error()), 409)
		return
	}
	err = utils.CheckPassword(user.Password)
	if err != nil {
		c.Error(T(err.Error()), 403)
		return
	}

	res, err := u.db.Exec(`
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3) RETURNING id
		`, user.Username, user.Email, user.Password)
	if err != nil {
		log.Print(err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Print(err)
		return
	}
	user.Id = int(id)
	c.JSON(200, user)
}
