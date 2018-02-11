package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gogap/logrus"
)

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		logrus.Infof("I got a %s!", r.Method)
		logrus.Infof("Body: %s", string(body))

		if string(body) != "" {
			w.WriteHeader(http.StatusOK)
			w.Write(body)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write("")
	})

	logrus.Fatal(http.ListenAndServe(":8080", nil))
}
