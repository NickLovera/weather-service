package contextKey

import (
	"context"
)

//This would be stored in a library so it could be used anywhere this package simulates that

const UserAgentStringKey = string("user-agent")

func EmbedUserAgent(c context.Context, userAgent string) context.Context {
	return context.WithValue(c, UserAgentStringKey, userAgent)
}

func GetUserAgent(c context.Context) string {
	userAgent, _ := c.Value(UserAgentStringKey).(string)
	return userAgent
}
