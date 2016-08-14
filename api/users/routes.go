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

// Searches for an user and checks his password.
// Returns the user with his id and JWT token.
//
// Errors:
// "users.not_found"
func (u *login) Serve(c *iris.Context) {
	T := utils.GetT(c)
	candidate := &User{}
	err := c.ReadJSON(candidate)
	if err != nil {
		log.Print(err)
		return
	}

	user := &User{}
	user.Username = candidate.Username
	err = u.db.QueryRow(`
		SELECT id, password, email FROM users WHERE username = $1
		`, candidate.Username).Scan(&user.Id, &user.Password, &user.Email)
	if err == sql.ErrNoRows {
		c.Error(T("users.not_found"), 404)
		return
	}
	if err != nil {
		log.Print(err)
		return
	}

	err = utils.VerifyPassword([]byte(candidate.Password), []byte(user.Password))
	if err != nil {
		c.Error(T(err.Error()), 401)
		return
	}
	user.Token, err = utils.GetToken(user.Id)
	if err != nil {
		log.Print(err)
		return
	}

	// don't send password to the user
	user.Password = ""
	c.JSON(200, user)
}

// Creates a new user, checking his uniqueness and password strength.
// Returns the user with his id and JWT token.
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

	hash, err := utils.HashPassword([]byte(user.Password))
	if err != nil {
		log.Print(err)
		return
	}
	user.Password = string(hash)
	err = u.db.QueryRow (`
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3) RETURNING id
		`, user.Username, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		log.Print(err)
		return
	}
	user.Token, err = utils.GetToken(user.Id)
	if err != nil {
		log.Print(err)
		return
	}

	// don't send password to the user
	user.Password = ""
	c.JSON(200, user)
}
