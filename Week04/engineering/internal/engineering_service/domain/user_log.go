package domain

import "log"

type UserLog struct {

}

func NewUserLog() *UserLog {
	return &UserLog{}
}

func (r *UserLog) WriteLog()  {
	log.Println("write log")
}
