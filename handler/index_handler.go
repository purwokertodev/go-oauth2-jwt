package handler

import (
	"github.com/wuriyanto48/go-json-message/response"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	response.MessageWithJson(res, "Your api ready to use", http.StatusOK)
}
