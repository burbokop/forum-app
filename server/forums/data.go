package forums

import (
	"database/sql"
	"strconv"
	"strings"
)

type Forum struct {
	Id           int64
	Name         string   `json:"name"`
	TopicKeyword string   `json:"topicKeyword"`
	Users        []string `json:"users"`
}

type AddUserRequest struct {
	Name      string   `json:"name"`
	Interests []string `json:"interests"`
}

type DBInterface struct{ Db *sql.DB }

func NewDBInterface(db *sql.DB) *DBInterface { return &DBInterface{Db: db} }

func TrimEachElem(slice []string) []string {
	var result []string
	for _, v := range slice {
		result = append(result, strings.Trim(v, " "))
	}
	return result
}

func unique(slice []string) []string {
	keys := make(map[string]bool)
	result := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}

func (dbi *DBInterface) ListForums() ([]*Forum, error) {
	rows, err := dbi.Db.Query("SELECT id, name, topic_keyword, subscribed_users FROM forums LIMIT 200")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*Forum
	for rows.Next() {
		var forum Forum
		var users_string string
		if err := rows.Scan(&forum.Id, &forum.Name, &forum.TopicKeyword, &users_string); err != nil {
			return nil, err
		}
		if len(users_string) > 0 {
			forum.Users = TrimEachElem(strings.Split(users_string, ","))
		}
		result = append(result, &forum)
	}
	if result == nil {
		result = make([]*Forum, 0)
	}
	return result, nil
}

func (dbi *DBInterface) AddUser(r *AddUserRequest) error {
	var requests []string
	var forums, err = dbi.ListForums()
	if err != nil {
		return err
	}
	for _, interest := range r.Interests {
		for _, forum := range forums {
			if interest == forum.TopicKeyword {
				var new_data string
				if len(forum.Users) > 0 {
					new_data = strings.Join(unique(append(forum.Users, r.Name)), ",")
				} else {
					new_data = r.Name
				}
				requests = append(requests, "UPDATE forums SET subscribed_users = '"+new_data+"' WHERE id="+strconv.FormatInt(forum.Id, 10))
			}
		}
	}

	for _, r := range requests {
		_, err = dbi.Db.Exec(r)
		if err != nil {
			return err
		}
	}

	return nil
}
