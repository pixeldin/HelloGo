package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
)

var store *sessions.CookieStore

func init() {
	// cookie加密秘钥
	store = sessions.NewCookieStore([]byte("sessionkey"))
	//store.Options = &sessions.Options{
	//	Path: "/",
	//	//session expire time
	//	MaxAge:   20,
	//	HttpOnly: true,
	//}
}

func GetStore() *sessions.CookieStore {
	return store
}

func GetSession(c *gin.Context, key string) *sessions.Session {
	session, e := store.Get(c.Request, key)
	if session.IsNew {
		//session.Options = &sessions.Options{
		//	Path: "/",
		//	//session expire time
		//	MaxAge:   20,
		//	HttpOnly: true,
		//}
		//session.Save(c.Request, c.Writer)
		log.Debug("=====================New session=====================")
	} else {
		log.Debug("=====================Get from old session=====================")
	}
	Logging(ErrCheck("Get session", e))
	return session
}

//func SaveSessionKey(c *gin.Context, key string, value interface{})  {
//	store.Save(c.Request, c.Writer, session)
//}

func ErrCheck(formatString string, err error) string {
	if err != nil {
		return fmt.Sprintf("Error occured with <%s>, err: %s", formatString, err)
	}
	return ""
}

func Logging(value string) bool {
	if value != "" {
		//maybe import log tools
		log.Info(value)
		return true
	}
	return false
}

func LoggingErr(value string) bool {
	if value != "" {
		//maybe import log tools
		log.Error(value)
		return true
	}
	return false
}

func Infof(msg string, val ...interface{})  {
	log.Info(msg, val)
}
