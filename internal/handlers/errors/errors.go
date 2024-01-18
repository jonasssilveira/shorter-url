package errors

import (
	"encoding/json"
)

type DatabaseError struct {
	MsgError   string `json:"msg_error"`
	Cause      string `json:"cause"`
	StatusCode int    `json:"status_code"`
}

func NewDatabaseError(msgError, cause string, statusCode int) *DatabaseError {
	return &DatabaseError{
		MsgError:   msgError,
		Cause:      cause,
		StatusCode: statusCode,
	}
}

func (e *DatabaseError) Error() string {
	marshal, err := json.Marshal(e)

	if err != nil {
		return err.Error()
	}
	return string(marshal)
}
