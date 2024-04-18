package types

type DataOpts struct {
  Pop int `json:"pop"`
  Quant int `json:"quant"`
  Props []Prop `json:"props"`
}

type Prop struct {
  Name string `json:"name"`
  Length int `json:"length"`
  Type string `json:"type"`
}
