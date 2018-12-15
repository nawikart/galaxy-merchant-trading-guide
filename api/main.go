package api

import (
		"io/ioutil"
		"encoding/json"
		"net/http"
		"../intergalacticConversion"
		"../romanNumerals2number"
		"../number2romanNumerals"
		)


var js []byte
var data = make(map[string]string)
var err error
// define struct to handle result of convertion
// Value: result of convertion, use type interface{} to handle result in integer (number) or string (roman numerals)
// Valid: handle validation of result
type Result struct {
	Value interface{}
	Valid bool
}
var result Result

// a function to handle conversion between number and roman numeral
func Conversion(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// grab data post as parameter convertion
	body, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(body, &data) == nil{
		input := data["input"]
		mod := data["mod"]

		// check mod parameter to determine the type of conversion
		switch mod {
			case "romanNumerals2number":
				result.Value, result.Valid = romanNumerals2number.Convert(input)
			case "number2romanNumerals":
				result.Value, result.Valid = number2romanNumerals.Convert(input)
		}
			
		result = result
	}
	js, _ = json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// a function to handle the answer about intergalactic units
func Intergalactic(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	body, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(body, &data) == nil{		
		js = intergalacticConversion.Answer(data)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}