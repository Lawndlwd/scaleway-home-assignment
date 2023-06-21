package models

import "gorm.io/gorm"

type Server struct {
	gorm.Model
    Name   string      
    Type   ServerType   
    Status ServerStatus 
}

type ServerType int

const (
    Small ServerType = iota
    Medium
    Large
)

type ServerStatus int

const (
    Starting ServerStatus = iota
    Running
    Stopping
    Stopped
)