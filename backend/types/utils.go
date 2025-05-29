package types

import "encoding/json"

type SailsError struct {
	Operation string `json:"operation"` // The operation that caused the error
	Reason    string `json:"reason"`    // The reason for the error
	Code      int    `json:"code"`      // HTTP status code
}

func (e *SailsError) Error() string {
	// Let me explain, currently in wails there is no clear way to return a custom error type
	// You only have 2 options:
	// 1. Return a string, which is an actual error and will reject the promise in the frontend
	// 2. Only return 'success' objects, but have the error in the 'error' field
	//    => This requires every returned object to have an 'error' field
	// So I decided to implement the first option, but then convert the error to JSON
	// This way, the frontend can simply parse the JSON into a proper error object and voila, type safety is preserved
	return e.ToJson()
}

func (e *SailsError) ToJson() string {
	jsonData, _ := json.Marshal(e)
	return string(jsonData)
}
