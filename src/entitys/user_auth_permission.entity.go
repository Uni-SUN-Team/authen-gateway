package entitys

import "time"

type UserAuthPermission struct {
	UserId    int
	Token     string
	Iat       int
	Ext       int
	LastLogin time.Time
}
