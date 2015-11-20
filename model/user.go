package model

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func CreateUser(user *User) (err error) {
	_, err = db.QueryOne(user, `
        INSERT INTO users (name, type) VALUES (?name, ?type)
        RETURNING id
    `, user)
	return
}
func GetUser(id int64) (user *User, err error) {
	_, err = db.QueryOne(&user, `SELECT * FROM users WHERE id = ?`, id)
	return
}

func GetUsers() (users []User, err error) {
	_, err = db.Query(&users, `SELECT * FROM users`)
	return
}
