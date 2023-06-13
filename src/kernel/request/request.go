package request

import (
	"bytes"
	"fmt"
	contract2 "github.com/ArtisanCloud/PowerLibs/v3/logger/contract"
	"io"

	"net/http"
)

func LogRequest(logger contract2.LoggerInterface, request *http.Request) {
	var output bytes.Buffer
	// 前置中间件
	output.WriteString(fmt.Sprintf("%s %s ", request.Method, request.URL.String()))

	// print out request header
	output.Write([]byte("\r\nrequest header: { \r\n"))
	for k, vals := range request.Header {
		for _, v := range vals {
			output.Write([]byte(fmt.Sprintf("\t%s:%s\r\n", k, v)))
		}
	}
	output.Write([]byte("} \r\n"))

	// print out request body
	if request.Body != nil {

		output.Write([]byte("request body:"))
		var buf bytes.Buffer
		reader := io.TeeReader(request.Body, &buf)
		body, _ := io.ReadAll(reader)
		request.Body = io.NopCloser(&buf)
		output.Write(body)
	}

	logger.InfoF(output.String())
}
