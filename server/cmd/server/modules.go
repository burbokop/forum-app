//+build wireinject

package main

import (
	"github.com/burbokop/forum-app/server/forums"
	"github.com/google/wire"
)

// ComposeApiServer will create an instance of CharApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*ForumApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from channels package.
		forums.Providers,
		// Provide ChatApiServer instantiating the structure and injecting channels handler and port number.
		wire.Struct(new(ForumApiServer), "Port", "ListForumsHandler", "AddUserHandler"),
	)
	return nil, nil
}
