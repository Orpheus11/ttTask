package model

type Relationship struct {
	Id            int64  `json:"id"`
	User_id       int64  `json:user_id`
	Other_user_id int64  `json:other_user_id`
	State         string `json:"state"`
	Type          string `json:"type"`
}

func CreateRelationship(relationship *Relationship) (err error) {
	_, err = db.QueryOne(relationship, `
        INSERT INTO relationships (user_id,other_user_id,state, type) VALUES (?user_id,?other_user_id,?state,?type)
        RETURNING id
    `, relationship)
	return
}
func GetRelationship(user_id, other_user_id int64) (relationship Relationship, err error) {
	_, err = db.QueryOne(&relationship, `SELECT * FROM relationships WHERE user_id = ? and other_user_id = ?`, user_id, other_user_id)
	return
}

func GetRelationships() (relationships []Relationship, err error) {
	_, err = db.Query(&relationships, `SELECT * FROM relationships`)
	return
}

func GetRelationshipsByUserId(user_id int64) (relationships []Relationship, err error) {
	_, err = db.Query(&relationships, `SELECT * FROM relationships WHERE user_id=?`, user_id)

	return
}

func CreateUpdateRelationship(relationship *Relationship) (err error) {
	re, _ := GetRelationship(relationship.User_id, relationship.Other_user_id)
	if re.Id == 0 {
		_, err = db.QueryOne(relationship, `
	       INSERT INTO relationships (user_id,other_user_id,state, type) VALUES (?user_id,?other_user_id,?state,?type)
	       RETURNING id
	   `, relationship)
		return
	}
	_, err = db.QueryOne(relationship, `
	       UPDATE relationships SET state = ? WHERE user_id = ? and other_user_id = ?
	   `, relationship.State, relationship.User_id, relationship.Other_user_id)

	return
}
