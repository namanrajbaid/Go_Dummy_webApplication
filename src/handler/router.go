package handler

import (
	"github.com/gorilla/mux"
)

// var SubRouter *mux.Router
// var Router *mux.Router

// func init(){
// 	// TODO: evaluate reading this path from swagger
// 	basePath := "/dummy"
// 	// StrictSlash true ensures that /path/ redirects to /path and vice-versa
// 	Router = mux.NewRouter().StrictSlash(true)
// 	SubRouter = Router.PathPrefix(basePath).Subrouter()
// }

var Router = mux.NewRouter().StrictSlash(true)
var SubRouter = Router.PathPrefix("/dummy").Subrouter()
