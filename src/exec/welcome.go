package exec

import (
	"log"
	"net/http"
//	"html/template"
)

func Welcome(w http.ResponseWriter, r *http.Request, param []string) {	
	
	
	hello := []byte("Welcome World!!!")
//	hello := []byte(param[0])
	
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}

}
