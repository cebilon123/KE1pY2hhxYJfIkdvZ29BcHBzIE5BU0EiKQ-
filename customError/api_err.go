package customError

import "encoding/json"

// ApiErr represents error which should be returned from api
type ApiErr struct {
	Error string `json:"error"`
}

func NewApiErr(error string) *ApiErr {
	return &ApiErr{Error: error}
}

func (e *ApiErr) ToByteSlice() []byte {
	json, err := json.Marshal(e)

	if err != nil {
		return nil
	}

	return json
}
