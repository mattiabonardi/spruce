package models

type ExecutionContext struct {
	UserSession UserSession
}

type UserSession struct {
	SessionId string
	Username  string
}
