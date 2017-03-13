package handler

import(
	"net/http"
	"github.com/wuriyanto48/go-json-message/response"
)

func Index(res http.ResponseWriter, req *http.Request){
		response.MessageWithJson(res, "Your api ready to use", http.StatusOK)
}