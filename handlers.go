package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func delayHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Info().Printf("Incomming request %s from %s\n", r.URL, r.RemoteAddr)

	d := ps.ByName("delay")

	delay, err := strconv.Atoi(d)
	if err != nil {
		log.Error().Println("Error converting", d, "to int")
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error converting '%s' to int", d)
		return
	}

	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func statusCodeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Info().Printf("Incomming request %s from %s\n", r.URL, r.RemoteAddr)

	c := ps.ByName("code")

	statusCode, err := strconv.Atoi(c)
	if err != nil {
		log.Error().Printf("Error converting '%s' to int\n", c)
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error converting '%s' to int", c)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			log.Error().Println("recover:", err)
			w.WriteHeader(999)
		}
	}()
	w.WriteHeader(statusCode)

	if r.ContentLength > 0 {
		io.Copy(w, r.Body)
	}

}

func dropConnectionHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info().Printf("Incomming request %s from %s\n", r.URL, r.RemoteAddr)

	hj, ok := w.(http.Hijacker)
	if !ok {
		log.Error().Println("Connection could not be hijacked")
		w.WriteHeader(400)
		return
	}

	conn, _, err := hj.Hijack()
	if err != nil {
		log.Error().Println("Error hijacking connection", err)
		w.WriteHeader(400)
		return
	}

	conn.Close()
}

func showHeadersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info().Printf("Incomming request %s from %s\n", r.URL, r.RemoteAddr)

	log.Info().Printf("%+v\n", r.Header)

	fmt.Fprintf(w, "%+v\n", r.Header)
}
