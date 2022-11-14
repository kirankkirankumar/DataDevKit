// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Meta struct {
	Count int `json:"count"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Notification struct {
	ID   string    `json:"id"`
	Date time.Time `json:"date"`
	Type string    `json:"type"`
}

type Stat struct {
	ID        string `json:"id"`
	Views     int    `json:"views"`
	Likes     int    `json:"likes"`
	Retweets  int    `json:"retweets"`
	Responses int    `json:"responses"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type Tweet struct {
	ID     string    `json:"id"`
	Body   string    `json:"body"`
	Date   time.Time `json:"date"`
	Author *User     `json:"Author"`
	Stats  *Stat     `json:"Stats"`
}

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}
