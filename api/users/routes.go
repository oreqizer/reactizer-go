package users

import (
	"database/sql"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"

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
func (r *login) Serve(c iris.Context) {
	candidate := &User{}
	err := c.ReadJSON(candidate)
	if err != nil {
		golog.Error(err)
		return
	}

	user := &User{}
	user.Username = candidate.Username
	err = r.db.QueryRow(`
		SELECT id, password, email FROM users WHERE username = $1
	`, candidate.Username).Scan(&user.Id, &user.Password, &user.Email)
	if err == sql.ErrNoRows {
		utils.Error(c, notFound, 404)
		return
	}
	if err != nil {
		golog.Error(err)
		return
	}

	err = utils.VerifyPassword([]byte(candidate.Password), []byte(user.Password))
	if err != nil {
		utils.Error(c, err.Error(), 401)
		return
	}
	user.Token, err = utils.GetToken(user.Id)
	if err != nil {
		golog.Error(err)
		return
	}

	// don't send password to the user
	user.Password = ""
	c.JSON(user)
}

// Creates a new user, checking his uniqueness and password strength.
// Returns the user with his id and JWT token.
func (r *register) Serve(c iris.Context) {
	user := &User{}
	err := c.ReadJSON(user)
	if err != nil {
		golog.Error(err)
		return
	}

	err = checkUsername(user.Username, r.db)
	if err != nil {
		utils.Error(c, err.Error(), 409)
		return
	}
	err = checkEmail(user.Email, r.db)
	if err != nil {
		utils.Error(c, err.Error(), 409)
		return
	}
	err = utils.CheckPassword(user.Password)
	if err != nil {
		utils.Error(c, err.Error(), 403)
		return
	}

	hash, err := utils.HashPassword([]byte(user.Password))
	if err != nil {
		golog.Error(err)
		return
	}
	user.Password = string(hash)
	err = r.db.QueryRow(`
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3) RETURNING id
	`, user.Username, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		golog.Error(err)
		return
	}
	user.Token, err = utils.GetToken(user.Id)
	if err != nil {
		golog.Error(err)
		return
	}

	// don't send password to the user
	user.Password = ""
	c.JSON(user)
}
