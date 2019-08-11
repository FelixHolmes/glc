package session

import "github.com/jueyoq/glc/conn"

type Session interface {
	Get(key uint64) (conn *conn.Conn, ok bool)
	Set(key uint64, value *conn.Conn) error
}
