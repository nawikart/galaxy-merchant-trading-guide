package intergalacticConversion

import (
		"encoding/json"
		"../romanNumerals2number"
		"strings"
		"strconv"
		)

var js []byte

// struct to collect pairs of intergalatic unit name and roan symbol
type TestInputs struct{
	Unit, RomanS string
}

// struct to collect pairs of set intergalatic units, 
// item (such as: silver, gold, iron), and number of credit
type TestInputs2 struct{
	Units, Item, Credits string
}

// struct to collect of pairs of question and type
type Question struct{
	Q, Type string
}

// struct to handle result of convertion
type Result struct {
	Value interface{}
	Valid bool
}
var result Result
func Answer(data map[string]string) []byte{
	var testInputs []TestInputs
	var testInputs2 []TestInputs2
	var q Question
	result.Valid = true

	// grab post data and convert to byte
	// broken into each corresponding struct
	datatestInputs := []byte(data["testInputs"])
	datatestInputs2 := []byte(data["testInputs2"])
	dataQuestion := []byte(data["faq"])

	// define map to collect value of each unit name with it's roman symbol
	mUnit2romanS := make(map[string]string)

	// define map to collect value of each item name with it's number of credits
	mItem2Credits := make(map[string]float64)


	// parse string of json into corresponding Map
	if json.Unmarshal(datatestInputs, &testInputs) == nil{
		for _, e := range(testInputs){
			mUnit2romanS[strings.ToLower(e.Unit)] = e.RomanS
		}
	}

	// parse string of json into corresponding Map
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
				mItem2Credits[strings.ToLower(e.Item)] = float64(credits) / float64(number.(int64))
			}
		}
	}

	// parse string of json into corresponding Map
	if json.Unmarshal(dataQuestion, &q) == nil{
		if q.Type == "number"{ // this is for question like "how much is pish tegj glob glob?"
			result.Value, result.Valid = units2number(q.Q, mUnit2romanS)
		}else{
			// this is for question like "how many credits glob prok Gold?"
			// firstly split it become "glob" "prok" Gold"
			qSplited := strings.Split(q.Q, " ")
			// grab last index to get the last words --> "Gold"
			indItem := len(qSplited)-1
			var romanV interface{}
			var valid bool
		
			if len(qSplited) == 1 {
				romanV = 1
			}else{
				romanS := ""
				for i, e := range(qSplited){
					// concat the roman symbol base on unit name
					// except the last index, because the last index in Item (eg. silver, gold. iron)
					if i < indItem {

						// if not found the unit name in Map, the make this invalid
						if mUnit2romanS[e] == "" {
							result.Value, result.Valid = nil, false
							break
						}else{
							// collect the roman symbols
							romanS = romanS + mUnit2romanS[e]
						}
					}
				}
			
				romanV, valid = romanNumerals2number.Convert(romanS)
			}				

			if valid {
				result.Value, result.Valid = mItem2Credits[strings.ToLower(qSplited[indItem])] * float64(romanV.(int64)), true
			}else{
				result.Value, result.Valid = nil, false
			}
		}
	}
	
	js, _ = json.Marshal(result)
	return js
}

// a function for conversion units to number
// eg. pish tegj glob glob --> 42
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