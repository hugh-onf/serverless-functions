package utils

type JsonRpc struct {
	Version string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      int      `json:"id"`
}

func NewJsonRpcV2(method string, params []string) *JsonRpc {
	if params == nil {
		params = []string{}
	}
	return &JsonRpc{
		Version: "2.0",
		Id:      1,
		Method:  method,
		Params:  params,
	}
}
