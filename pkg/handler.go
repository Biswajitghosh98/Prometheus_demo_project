package pkg

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
)

type Server struct {
	NumRequestHandled prometheus.Counter
	Num200Requests    prometheus.Counter
	NumNon200Requests prometheus.Counter
	NumUniqueRequests prometheus.Counter
}

func (srv *Server) ReturnUniqueNumber(w http.ResponseWriter, r *http.Request) {
	uid := uuid.New()
	fmt.Printf("%v", uid)
	s := fmt.Sprintf("your new request ID is : %v", uid)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(s))
	srv.NumUniqueRequests.Inc()
}
func (srv *Server) ReturnIfOdd(w http.ResponseWriter, r *http.Request) {
	val := r.Header.Get("number")
	num, _ := strconv.Atoi(val)
	fmt.Printf("%d", num)
	if num%2 == 1 {
		w.WriteHeader(http.StatusAccepted)
		s := fmt.Sprintf("The number given is odd")
		w.Write([]byte(s))
		return
	}
	w.WriteHeader(http.StatusInternalServerError)

}
