package server

// ID's of service results
const (
	OK               = 0x00
	NOT_AN_EMAIL     = 0x01
	INVALID_NAME     = 0x02
	INVALID_FEEDBACK = 0x03
	DB_UNAVAILABLE   = 0x05
	CANT_WRITE_TO_DB = 0x06
)
