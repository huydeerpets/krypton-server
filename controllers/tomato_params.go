package controllers

import "time"

type CreateTomatoParams struct {
	Uid      string    `json:"uid"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Desc     string    `json:"desc"`
	Finished bool      `json:"finished"`
}

type UpdateTomatoParams struct {
	ID       string    `json:"id"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Desc     string    `json:"desc"`
	Finished bool      `json:"finished"`
}
