package common

import (
	"github.com/gin-gonic/gin"
)

const (
	StatusSuccessCode       = 10000 // success
	StatusFailureCode       = 10001 // failed
	StatusErrorCode         = 10002 // internal error
	StatusRequestedSelfCode = 10003 // requested self
)
const (
	StatusSuccess       = "success"
	StatusError         = "error"
	StatusUnauthorized  = "unauthorized"
	StatusFailure       = "failure"
	StatusRequestedSelf = "requested_self"
)

type StandardMessage struct {
	Code    int               `json:"code" binding:"required"`
	Msg     string            `json:"msg"`
	Status  string            `json:"status" binding:"required"`
	Links   map[string]string `json:"links,omitempty"`
	Content interface{}       `json:"content"`
}

func FormatMessageBody(message StandardMessage) gin.H {
	return gin.H{
		"code":    message.Code,
		"status":  message.Status,
		"msg":     message.Msg,
		"links":   message.Links,
		"content": message.Content,
	}
}

func StatusBadRequestMessage(msg string) gin.H {
	return FormatMessageBody(StandardMessage{Code: StatusFailureCode, Status: StatusFailure, Msg: msg})
}
func StatusInternalServerErrorMessage(msg string) gin.H {
	return FormatMessageBody(StandardMessage{Code: StatusErrorCode, Status: StatusError, Msg: msg})
}
func StatusUnauthorizedMessage(msg string) gin.H {
	return FormatMessageBody(StandardMessage{Code: StatusFailureCode, Status: StatusUnauthorized, Msg: msg})
}
func StatusOKMessage(content map[string]any, msg string) gin.H {
	return FormatMessageBody(StandardMessage{
		Code:    StatusSuccessCode,
		Status:  StatusSuccess,
		Msg:     msg,
		Content: content,
	})
}
func StatusRequestedSelfMessage(msg string) gin.H {
	return FormatMessageBody(StandardMessage{Code: StatusRequestedSelfCode, Status: StatusRequestedSelf, Msg: msg})
}
