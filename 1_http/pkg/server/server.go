package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var indexPage = `
<!DOCTYPE html>
<html>
	<body>
		<h1 style="text-align:center;"> User Database </h1>
		<p style="text-align:center;"> 
		Welcome to the user database.
		</p>
	</body>
</html>
`

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint8  `json:"age"`
}

type userinfo struct {
	email string
	age   uint8
}

// Server is an HTTP server.
type Server struct {
	users map[string]userinfo
}

// New returns a new server.
func New() *Server {
	return &Server{
		users: make(map[string]userinfo),
	}
}

// HandleIndex handles the index (root) route.
func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html") // Encoding standards
	w.Write([]byte(indexPage))
	w.WriteHeader(http.StatusOK)
}

// HandleCreateUser handles the `/user/create` POST/PUT route.
func (s *Server) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost, http.MethodPut:
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType) // HTTP 415
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var user user
		err = json.Unmarshal(body, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Create user: %v", user.Name)
		s.users[user.Name] = userinfo{
			email: user.Email,
			age:   user.Age,
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// HandleUser handles the HTTP Get method on `/users/` path
func (s *Server) HandleUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	u, ok := s.users[name]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		ret := user{
			Name:  name,
			Email: u.email,
			Age:   u.age,
		}
		msg, err := json.Marshal(ret)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("Found user: %v", name)
		w.Header().Add("Content-Type", "application/json")
		w.Write(msg)
	case http.MethodPatch:
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType) // HTTP 415
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var user user
		err = json.Unmarshal(body, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Update user: %v", name)

		userinfo := s.users[name]
		if user.Age != 0 {
			userinfo.age = user.Age
		}
		if user.Email != "" {
			userinfo.email = user.Email
		}
		s.users[name] = userinfo
	case http.MethodDelete:
		_, ok := s.users[name]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Printf("Delete user: %v", name)
		delete(s.users, name)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
