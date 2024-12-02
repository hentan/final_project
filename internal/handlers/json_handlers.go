package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/hentan/final_project/internal/logger"
	_ "github.com/hentan/final_project/internal/repository"
)

type JSONResponce struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (app *Application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (app *Application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("поддерживается только 1 json значение в body")
	}

	return nil
}

func (app *Application) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	newLogger := logger.GetLogger()

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponce
	payload.Error = true
	payload.Message = err.Error()
	newLogger.Error(payload.Message)
	sErr := fmt.Sprint(err)
	newLogger.Info("строка 75")
	er := app.KafkaClient.SendMessage(sErr)
	if er != nil {
		sEr := fmt.Sprint(er)
		newLogger.Error(sEr)
		return err
	}
	newLogger.Info("строка 82")
	return app.writeJSON(w, statusCode, payload)
}
