package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	pgx "github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Gender   string    `json:"gender"`
	Dob      time.Time `json:"dob"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
}

type Server struct {
	db *pgx.Pool
}

func (server *Server) helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
func (server *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	results, err := server.db.Query(context.Background(), "SELECT * FROM USERS WHERE dob > DATE('1990-01-01')")
	if err != nil {
		log.Fatal(err)
	}
	var users []*User
	for results.Next() {
		var u User
		err = results.Scan(&u.Id, &u.Username, &u.Name, &u.Gender, &u.Dob, &u.Email, &u.Phone)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &u)
	}
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Keep-Alive", "timeout=5")
	w.Write(jsonData)
}

func main() {
	conn := "postgresql://admin:admin@127.0.0.1:5432/db?sslmode=disable"
	db, err := pgx.New(context.Background(), conn)
	if err != nil {
		log.Fatal(err)
	}
	server := Server{db}
	http.HandleFunc("/", server.userHandler)
	fmt.Println("Server is running on port 8003...")
	http.ListenAndServe(":8003", nil)
}

func (server *Server) getUsers() (interface{}, error) {
	rows, err := server.db.Query(context.Background(), "SELECT * FROM USERS WHERE dob > DATE('1990-01-01')")
	return &rows, err
}
