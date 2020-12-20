package forums

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/burbokop/forum-app/server/tools"
)

// Channels HTTP handler.
type HttpVmListHandlerFunc http.HandlerFunc
type HttpConnectDiskHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpVmListHandler(dbi *DBInterface) HttpVmListHandlerFunc {
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

func HttpConnectDiskHandler(dbi *DBInterface) HttpConnectDiskHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Println("Disk connecting request recieved.")
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
