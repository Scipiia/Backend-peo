package get_json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/http"
	"vue-golang/internal/storage"
)

type Importjson interface {
	ImportJSON(data storage.ImportJSON) error
}

func New(log *slog.Logger, importJSON Importjson) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.importjson.new"

		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		log.Info("SDSDD", body)

		var person storage.ImportJSON
		err = json.Unmarshal(body, &person)
		if err != nil {
			http.Error(w, "Error decoding JSON1", http.StatusBadRequest)
			return
		}

		log.Info("PERSON", person)
		err = importJSON.ImportJSON(person)
		if err != nil {
			http.Error(w, "Error decoding JSON2", http.StatusBadRequest)
			return
		}

		fmt.Printf("Получен пользователь: %+v\n", person)

		// Отправляем ответ клиенту
		response := map[string]string{"status": "success", "message": "Данные получены"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
