package http_utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func DecodeRequestBodyToJson(body io.ReadCloser, target any) error {
	err := json.NewDecoder(body).Decode(target)
	if err != nil {
		return err
	}

	return nil
}

func SetErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)

	if err == nil {
		write(w, nil)
		return
	}

	je := formatJSONError(err.Error())
	write(w, je)
}

func write(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		log.Fatalln(err)
	}
}

func formatJSONError(message string) []byte {
	appError := struct {
		Message string `json:"message"`
	}{
		message,
	}

	response, err := json.Marshal(appError)

	if err != nil {
		return []byte(err.Error())
	}

	return response
}
