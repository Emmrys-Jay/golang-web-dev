package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type codes struct {
	Code    int    `json:"Code"`
	Descrip string `json:"Descrip"`
}

func main() {
	var rcvd = `[{"Code":200,"Descrip":"StatusOK"},{"Code":301,"Descrip":"StatusMovedPermanently"},{"Code":302,"Descrip":"StatusFound"},{"Code":303,"Descrip":"StatusSeeOther"},{"Code":307,"Descrip":"StatusTemporaryRedirect"},{"Code":400,"Descrip":"StatusBadRequest"},{"Code":401,"Descrip":"StatusUnauthorized"},{"Code":402,"Descrip":"StatusPaymentRequired"},{"Code":403,"Descrip":"StatusForbidden"},{"Code":404,"Descrip":"StatusNotFound"},{"Code":405,"Descrip":"StatusMethodNotAllowed"},{"Code":418,"Descrip":"StatusTeapot"},{"Code":500,"Descrip":"StatusInternalServerError"}] `

	var data []codes
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Println("could not unmarshal json", err)
	}
	fmt.Println(data)
	//fmt.Println(data)
	//fmt.Println(data)
}
