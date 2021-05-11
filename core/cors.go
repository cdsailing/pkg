package core

import "github.com/emicklei/go-restful/v3"

func Cors(container restful.Container) restful.CrossOriginResourceSharing {
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: true,
		Container:      &container,
	}
	return cors
}
