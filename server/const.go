package server

import (
	"net/http"
)

const (
	ServerName = "FindWedding"
	Empty      = 0
)

type MethodFunc func(w http.ResponseWriter, r *http.Request)
