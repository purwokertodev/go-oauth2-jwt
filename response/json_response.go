package response

import(
	"net/http"
	"encoding/json"
)

func MessageWithJson(res http.ResponseWriter, message string, httpCode int){
	msg, _ := json.Marshal(message)
	res.Header().Set("Content-Type", "application-json; charset=utf-8")
	res.WriteHeader(httpCode)
	res.Write(msg)
}

func ResponseWithJson(res http.ResponseWriter, json []byte, httpCode int){
	res.Header().Set("Content-Type", "application-json; charset=utf-8")
	res.WriteHeader(httpCode)
	res.Write(json)
}

func JsonResponse(res http.ResponseWriter, resp interface{}, httpCode int){
	msg, _ := json.Marshal(resp)
	res.Header().Set("Content-Type", "application-json; charset=utf-8")
	res.WriteHeader(httpCode)
	res.Write(msg)
}