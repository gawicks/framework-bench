package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pgx "github.com/jackc/pgx/v5/pgxpool"
	"github.com/valyala/fasthttp"
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

func (server *Server) userHandler(ctx *fasthttp.RequestCtx) {
	results, err := server.db.Query(context.Background(), "SELECT * FROM USERS WHERE dob > '1990-01-01'")
	if err != nil {
		log.Fatal(err)
	}
	defer results.Close()

	var users []User
	for results.Next() {
		var u User
		err = results.Scan(&u.Id, &u.Username, &u.Name, &u.Gender, &u.Dob, &u.Email, &u.Phone)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	ctx.SetBody(jsonData)
}

func main() {
	conn := "user=admin password=admin host=localhost port=5432 dbname=db sslmode=disable pool_max_conns=40"
	db, err := pgx.New(context.Background(), conn)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	server := Server{db}
	if err := fasthttp.ListenAndServe(":8003", server.userHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}
