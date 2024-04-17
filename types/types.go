package types

type DataOpts struct {
  Pop int `json:"pop"`
  Quant int `json:"quant"`
  Props []map[string]string `json:"props"`
}
