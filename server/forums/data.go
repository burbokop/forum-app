package forums

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type Forum struct {
	Id             int64
	Name           string `json:"name"`
	TopicKeyword   string `json:"topicKeyword"`
	Users string[]  `json:"users"`
}

type AddUserRequest struct {
	Name string `json:"name"`
	Interests string[] `json:"interests"`
}

type DBInterface struct { Db *sql.DB }
func NewDBInterface(db *sql.DB) *DBInterface { return &DBInterface{Db: db} }

func Contains(s []string, elem string) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}
	return false
}

func (s *DBInterface) ListForums() ([]*Forum, error) {
	rows, err := s.Db.Query("SELECT id, name, topic_keyword, subscribed_users FROM forums LIMIT 200")
	if err != nil {
		return nil, err
	}	
	defer rows.Close()

	var result []*Forum
	for rows.Next() {
		var forum Forum
		if err := rows.Scan(&forum.Id, &forum.Name, &forum.TopicKeyword, &forum.Users); err != nil {
			return nil, err
		}
		result = append(result, &c)
	}
	if result == nil {
		result = make([]*VirtualMachine, 0)
	}
	return result, nil
}


func (s *VMStorage) AddUser(request *AddUserRequest) error {
	var forums = ListForums();

	request.

	rows, err := s.Db.Query("SELECT connected_discs FROM virtual_machines WHERE id=$1", arg.VmId)
	if err != nil {
		return err
	}
	var cd string
	if rows.Next() {
		err = rows.Scan(&cd)
		if err != nil {
			return err
		}
	} else {
		return &VMNotFoundError{}
	}
	var parsed_cd = removeDuplicateValues(ParseDiscListString(cd))
	if ContainsInSlice(parsed_cd, arg.DiskId) {
		return nil
	}

	parsed_cd = append(parsed_cd, arg.DiskId)
	var new_data string
	for i, elem := range parsed_cd {
		new_data += strconv.FormatInt(elem, 10)
		if i < len(parsed_cd)-1 {
			new_data += ", "
		}
	}
	var request = "UPDATE virtual_machines SET connected_discs = '" + new_data + "' WHERE id=" + strconv.FormatInt(arg.VmId, 10)
	_, err = s.Db.Exec(request)

	return err
}
