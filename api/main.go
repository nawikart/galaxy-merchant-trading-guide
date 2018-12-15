package api

import (
		"io/ioutil"
		"encoding/json"
		"net/http"
		"../romanNumerals2number"
		"../number2romanNumerals"
		// "fmt"
		"strings"
		"strconv"
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


type TestInputs struct{
	Unit, RomanS string
}
type TestInputs2 struct{
	Units, Item, Credits string
}
type Question struct{
	Q, Type string
}

func Answer(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var testInputs []TestInputs
	var testInputs2 []TestInputs2
	var q Question
	result.Valid = true

	body, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(body, &data) == nil{
		
		datatestInputs := []byte(data["testInputs"])
		datatestInputs2 := []byte(data["testInputs2"])
		dataQuestion := []byte(data["faq"])
		mUnit2romanS := make(map[string]string)
		mUnit2romanS2 := make(map[string]float64)

		if json.Unmarshal(datatestInputs, &testInputs) == nil{
			for _, e := range(testInputs){
				mUnit2romanS[strings.ToLower(e.Unit)] = e.RomanS
			}
		}

		if json.Unmarshal(datatestInputs2, &testInputs2) == nil{
			for _, e := range(testInputs2){
				credits, err := strconv.Atoi(e.Credits)	
				if err != nil {
					result.Value, result.Valid = nil, false
					break
				}
				number, validConvert := units2number(e.Units, mUnit2romanS)
				if !validConvert {
					result.Value, result.Valid = nil, false
					break					
				}else{
					mUnit2romanS2[strings.ToLower(e.Item)] = float64(credits) / float64(number.(int64))
				}
			}
		}

		if json.Unmarshal(dataQuestion, &q) == nil{
			
			if q.Type == "number"{
				result.Value, result.Valid = units2number(q.Q, mUnit2romanS)
			}else{
				qSplited := strings.Split(q.Q, " ")
				indItem := len(qSplited)-1
				var romanV interface{}
				var valid bool
			
				if len(qSplited) == 1 {
					romanV = 1
				}else{
					romanS := ""
					for i, e := range(qSplited){
						if i < indItem {
							if mUnit2romanS[e] == "" {
								result.Value, result.Valid = nil, false
								break
							}else{
								romanS = romanS + mUnit2romanS[e]
							}
						}
					}
				
					romanV, valid = romanNumerals2number.Convert(romanS)
				}				

				if valid {
					result.Value, result.Valid = mUnit2romanS2[strings.ToLower(qSplited[indItem])] * float64(romanV.(int64)), true
				}else{
					result.Value, result.Valid = nil, false
				}
			}
		}		
	}
	js, _ = json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func units2number(units string, mUnit2romanS map[string]string) (interface{}, bool){
	qSplited := strings.Split(units, " ")
	romanS := ""
	valid := true
	for _, e := range(qSplited){
		if mUnit2romanS[strings.ToLower(e)] == "" {
			valid = false
			break
		}else{
			romanS = romanS + mUnit2romanS[strings.ToLower(e)]
		}
	}
	if valid {
		return romanNumerals2number.Convert(romanS)
	}else{
		return nil, false
	}
}