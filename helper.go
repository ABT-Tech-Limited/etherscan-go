package etherscan

import "encoding/json"

func ToMap(req any) map[string]string {
	tmp, _ := json.Marshal(req)
	m := make(map[string]string)
	_ = json.Unmarshal(tmp, &m)
	return m
}
