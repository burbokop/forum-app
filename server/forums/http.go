package forums

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/burbokop/forum-app/server/tools"
)

type HttpListForumsHandlerFunc http.HandlerFunc
type HttpAddUserHandlerFunc http.HandlerFunc

func HttpListForumsHandler(dbi *DBInterface) HttpListForumsHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			res, err := dbi.ListForums()
			if err != nil {
				log.Printf("Error making query to the db: %s", err)
				tools.WriteJsonInternalError(rw)
				return
			}
			tools.WriteJsonOk(rw, res)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func HttpAddUserHandler(dbi *DBInterface) HttpAddUserHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var c AddUserRequest
			if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
				log.Printf("Error decoding request input: %s", err)
				tools.WriteJsonBadRequest(rw, "bad JSON payload")
				return
			}
			err := dbi.AddUser(&c)
			if err == nil {
				res, err := dbi.ListForums()
				if err == nil {
					tools.WriteJsonOk(rw, res)
				} else {
					log.Println(err)
				}
			} else {
				log.Printf("Error inserting record: %s", err)
				tools.WriteJsonInternalError(rw)
			}
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
