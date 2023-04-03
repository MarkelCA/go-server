package router

import (
    "net/http"
)

type httpMethod string

const (
	mGET httpMethod = http.MethodGet
	mPOST           = http.MethodPost
	mPUT            = http.MethodPut
	mDELETE         = http.MethodDelete
)

var strToMethod = map[string]httpMethod{
    http.MethodGet    : mGET,
    http.MethodDelete : mDELETE,
    http.MethodPost   : mPOST,
    http.MethodPut    : mPUT,
}


// Receives a url string and if it ends 
// with / it removes it to match the original route
func removeTrailingSlash(url string) string {
    lastURLChar := url[len(url)-1:]
    if lastURLChar == "/" {
        url = url[:len(url)-1]
    }

    return url

}

