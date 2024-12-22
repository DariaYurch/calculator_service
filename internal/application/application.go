package application

import (
	"github.com/DariaYurch/calculator_service/pkg/calculator"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type InputData struct{
	Expression string `json:"expression"`
}
type ResultData struct {

	Result float64 `json:"result"`
}

type Error struct{
	Error string `json:"error"`
}

func CalculatorHandler(w http.ResponseWriter, r *http.Request){
	var inp InputData
	defer r.Body.Close()
	if r.Method == "POST"{
		body, err := ioutil.ReadAll(r.Body)
		if err != nil{
			log.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError) //500
			json.NewEncoder(w).Encode(Error{Error: err.Error()})
			return
		}
		err = json.Unmarshal(body, &inp)
		if err != nil{
			log.Printf("Can`t unmarshal JSON: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError) //500
			json.NewEncoder(w).Encode(Error{Error: err.Error()})
			return
		}
		data, err := calculator.Calc(inp.Expression)
		if err != nil{
			log.Printf("Can`t calculate expression: %s", err.Error())
			w.WriteHeader(http.StatusUnprocessableEntity) //422
			json.NewEncoder(w).Encode(Error{Error: err.Error()})
			return
		}
		response := ResultData{Result: data}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil{
			log.Printf("Error encoding JSON: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError) //500
			json.NewEncoder(w).Encode(Error{Error: err.Error()})
			return
		}

	}else{
		log.Println("Method isn`t supported")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Error: "Method isn`t supported"})
		return
	}
}

func RunServer() error{
	http.HandleFunc("/api/v1/calculate", CalculatorHandler)
	log.Printf("Server is running")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error with starting the server:", err)
		return nil
	} else {
		return err
	}
}

