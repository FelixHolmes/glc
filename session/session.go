package session

import "github.com/jueyoq/glc/conn"

type Session interface {
	Get(key interface{}) (conn *conn.Conn, ok bool)
	Set(key interface{}, value *conn.Conn) error
}
