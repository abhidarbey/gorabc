package server

import (
	"gorabc/pkg/controllers/routes"
)

func mapUrls() {
	routes.Ping(router)
	routes.Auth(router)
	routes.Users(router)
	routes.Organization(router)
	routes.Department(router)
	routes.Role(router)
}
