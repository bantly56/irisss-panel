package api

import (
	"net/http"
)

func Logs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("Irisss Panel log system is under development\n"))
}
