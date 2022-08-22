package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/teris-io/shortid"
)

func HandleUrl(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write([]byte(fmt.Sprintf(`{"short_url": "http://0.0.0.0:8080/%s", "original_url": "%s"}`, shortid.MustGenerate(), data["url"])))
	if err != nil {
		w.WriteHeader(500)
		return
	}
}
