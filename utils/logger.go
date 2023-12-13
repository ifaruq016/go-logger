package utils

import (
	"fmt"
	"log"
	"encoding/json"
	"sync"
	"github.com/gin-gonic/gin"
)

type DataInfo struct {
	QrId     		string
	Resp     		map[string]interface{}
	Data     		string
	State    		string
	Oneline  		bool
}

type StroreRequestInfo struct{
	ApiId			string
	RequestBody		string
	RequestFile		string
}

type StorageRequestInfo = struct {
	sync.RWMutex
	data map[string]StroreRequestInfo
}

var storage = StorageRequestInfo{
	data : make(map[string]StroreRequestInfo),
}

func Store(key string, data StroreRequestInfo) {
	storage.Lock()
	defer storage.Unlock()

	storage.data[key] = data
}

func Get(key string) (StroreRequestInfo, bool) {
	storage.RLock()
	defer storage.RUnlock()

	data, ok := storage.data[key]
	return data, ok
}

func Delete(key string) {
	storage.Lock()
	defer storage.Unlock()

	delete(storage.data, key)
}

func LogInfo(ctx *gin.Context, apiId string, dataInfo DataInfo){
	defer func() {
		if rec := recover(); rec != nil {
			log.Println(fmt.Sprintf("%s|%s", apiId, "ERROR_INSERT_LOG"))
		}
	}()

	var logMessage string

	apiState := fmt.Sprintf("%s|%s", apiId, dataInfo.State)

	if dataInfo.Oneline {
		if dataInfo.Data != "" {
			logMessage = fmt.Sprintf("%s|%v", apiId, dataInfo.Data)
		}
	} else if dataInfo.State == "REQUEST_START" || dataInfo.State == "RESPONSE_END" {
		path := ctx.Request.URL.Path
		serverIP := ctx.Request.Host
		fullPath := fmt.Sprintf("%s%s", serverIP, path)
		clientIP := ctx.ClientIP()
		headers := ctx.Request.Header

		if dataInfo.State == "REQUEST_START" {

			storedData, found := Get(apiId)
			if found {
				var payload string
				var files string
				if len(storedData.RequestBody) > 0 {
					payload = fmt.Sprintf("payload:%v", storedData.RequestBody)
				}
				
				if len(storedData.RequestFile) > 0 {
					files = fmt.Sprintf("files:%v", storedData.RequestFile)
				}

				logMessage = fmt.Sprintf("%s|%s|%s|headers:%v|%v|%v|response:%v", apiId, clientIP, fullPath, headers, payload, files, "{}" )
			}
		} else if dataInfo.State == "RESPONSE_END" {
			response := "{}"
			if len(dataInfo.Resp) > 0 {
				dataString, err := convertMapToString(dataInfo.Resp)
				if err == nil {
					response = dataString
				}
			}
			logMessage = fmt.Sprintf("%s|%s|%s|headers:%v|response:%v", apiId, clientIP, fullPath, headers, response )
		}

	} else {
		data := "EMPTY_DATA_LOGGER"
		if len(dataInfo.Data) > 0 {
			data = dataInfo.Data
		}
		logMessage = fmt.Sprintf("%s|%v", apiId, data)
	}

	if dataInfo.State != "" {
		log.Println(apiState)
		log.Println(logMessage)
	} else {
		log.Println(logMessage)
	}
}

func RequestStore(ctx *gin.Context, apiId string) {
	var jsonData map[string]interface{}
	var jsonDataFile map[string]interface{}
	var jsonString string
	var jsonFileString string
	// var dataType string

	contentType := ctx.ContentType()
	if contentType == "application/json" {
		ctx.ShouldBindJSON(&jsonData)
	} else if contentType == "multipart/form-data" {
		jsonData = make(map[string]interface{})
		jsonDataFile = make(map[string]interface{})

		// Use c.Request.MultipartForm to access form data
		form, _ := ctx.MultipartForm()

		for key, values := range form.Value {
			if len(values) > 0 {
				jsonData[key] = values[0]
			}
		}

		for key, fileHeaders := range form.File {
			for _, fileHeader := range fileHeaders {
				// Extract only the filename without reading the entire file content
				jsonDataFile[key] = fileHeader.Filename
			}
		}
	} else {
		queryString := ctx.Request.URL.RawQuery
		data, _ := queryToJSON(queryString)
		jsonData = data
	}

	if len(jsonData) > 0 {
		jsonByte, _ := json.Marshal(jsonData)
		jsonString = string(jsonByte)
	}
	

	if len(jsonDataFile) > 0 {
		jsonFileByte, _ := json.Marshal(jsonDataFile)
		jsonFileString = string(jsonFileByte)
	}

	reqInfo := StroreRequestInfo{
        ApiId			:apiId,
        RequestBody		:jsonString,
        RequestFile		:jsonFileString,
    }
    Store(apiId, reqInfo)
}