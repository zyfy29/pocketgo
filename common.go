package pocket

import "encoding/json"

func JsonMarshal(v any) string {
	res, _ := json.Marshal(v)
	return string(res)
}

func JsonMarshalIndent(v any) string {
	res, _ := json.MarshalIndent(v, "", "  ")
	return string(res)
}
