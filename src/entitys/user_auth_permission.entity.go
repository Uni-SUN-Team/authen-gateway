package entitys

type UserAuthPermission struct {
	UserId       int
	TokenVersion int
	Iat          float64
	Ext          float64
}
