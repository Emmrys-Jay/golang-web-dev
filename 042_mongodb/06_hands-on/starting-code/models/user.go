package models

type UserVals struct {
	Name   string
	Gender string
	Age    int
	Id     string
}

type Users map[string]UserVals
