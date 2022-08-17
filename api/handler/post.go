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

	json.Unmarshal(b, &data)

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"short_url": "http://0.0.0.0:8080/%s", "original_url": "%s"}`, shortid.MustGenerate(), data["url"])))
	w.WriteHeader(201)
}
