package Server

import (
	"encoding/json"
	"github.com/DaniilOr/dbManipulations/src/models"
	"github.com/DaniilOr/dbManipulations/src/service"
	"log"
	"net/http"
	"strconv"
)
type Server struct {
	service service.ServiceInterface
	mux *http.ServeMux
}
func NewServer(service service.ServiceInterface, mux *http.ServeMux) *Server {
	return &Server{service: service, mux: mux}
}

func (s *Server) Init() {
	s.mux.HandleFunc("/getCards", s.getCards)
	s.mux.HandleFunc("/getTransactions", s.getTransactions)
	s.mux.HandleFunc("/getMostSpent", s.getMostSpent)
	s.mux.HandleFunc("/getMostVisited", s.getMostVisited)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
func(s*Server) getCards(w http.ResponseWriter, r *http.Request){
	suid := r.URL.Query().Get("uid")
	if suid != ""{
		uid, err := strconv.ParseInt(suid, 10, 64)
		if err != nil{
			response := models.Result{Result: "Error", ErrorDescription: "Bad uid"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		cards, err := s.service.GetCards(uid)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
		}
		if len(cards) == 0{
			response := models.Result{Result: "Error", ErrorDescription: "No results"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		respBody, err := json.Marshal(cards)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)

	} else {
		response := models.Result{Result: "Error", ErrorDescription: "No uid"}
		respBody, err := json.Marshal(response)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
	}
}

func (s*Server)  getTransactions(w http.ResponseWriter, r *http.Request){
	scid := r.URL.Query().Get("cid")
	if scid !=""{
		cid, err := strconv.ParseInt(scid, 10, 64)
		if err != nil{
			response := models.Result{Result: "Error", ErrorDescription: "Bad cid"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		transactions, err := s.service.GetTransactions(cid)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(transactions) == 0{
			response := models.Result{Result: "Error", ErrorDescription: "No results"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		respBody, err := json.Marshal(transactions)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
	} else {
		response := models.Result{Result: "Error", ErrorDescription: "No cid"}
		respBody, err := json.Marshal(response)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
		return
	}
}
func(s*Server) getMostSpent(w http.ResponseWriter, r*http.Request){
	scid := r.URL.Query().Get("cid")
	if scid !=""{
		cid, err := strconv.ParseInt(scid, 10, 64)
		if err != nil{
			response := models.Result{Result: "Error", ErrorDescription: "Bad cid"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		mcc, value, err := s.service.GetMostSpent(cid)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if mcc == ""{
			response := models.Result{Result: "Error", ErrorDescription: "No such card"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		response := models.MostSpentDTO{Mcc: mcc , Value: value}
		respBody, err := json.Marshal(response)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
	} else {
		response := models.Result{Result: "Error", ErrorDescription: "No cid"}
		respBody, err := json.Marshal(response)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
		return
	}
}

func(s*Server) getMostVisited(w http.ResponseWriter, r*http.Request){
	scid := r.URL.Query().Get("cid")
	if scid !=""{
		cid, err := strconv.ParseInt(scid, 10, 64)
		if err != nil{
			response := models.Result{Result: "Error", ErrorDescription: "Bad cid"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		mcc, value, err := s.service.GetMostVisited(cid)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if mcc == ""{
			response := models.Result{Result: "Error", ErrorDescription: "No such card"}
			respBody, err := json.Marshal(response)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			makeResponse(respBody, w, r)
			return
		}
		response := models.MostSpentDTO{Mcc: mcc , Value: value}
		respBody, err := json.Marshal(response)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
	} else {
		response := models.Result{Result: "Error", ErrorDescription: "No cid"}
		respBody, err := json.Marshal(response)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		makeResponse(respBody, w, r)
		return
	}
}

func makeResponse(respBody []byte, w http.ResponseWriter, r*http.Request) {
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write(respBody)
	if err != nil {
		log.Println(err)
	}
}
