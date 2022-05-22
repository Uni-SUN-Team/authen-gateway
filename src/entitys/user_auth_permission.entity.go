package entitys

import "time"

type UserAuthPermission struct {
	UserId       int
	TokenVersion string
	Iat          float64
	Ext          float64
	LastLogin    time.Time
}
