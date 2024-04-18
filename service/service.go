package service

import (
	"math"
	"math/rand"

	"github.com/Abdirahman04/jama/types"
)

func newStr(ln int) string {
  chars := [26]byte{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}

  var str []byte

  for i := 0;i < ln;i++ {
    str = append(str, chars[rand.Intn(26)])
  }

  return string(str)
}

func newInt(ln int) int {
  return rand.Intn(int(math.Pow(10, float64(ln))))
}

func newBool() bool {
  bools := []bool{true,false}
  return bools[rand.Intn(2)]
}

func newEmail(ln int) string {
  return newStr(ln) + "@gmail.com"
}

func GetSingleObj(props []types.Prop) map[string]interface{} {
  newProps := make(map[string]interface{})

  for _, v := range props {
    if v.Type == "string" {
      newProps[v.Name] = newStr(v.Length)
    } else if v.Type == "int" {
      newProps[v.Name] = newInt(v.Length)
    } else if v.Type == "bool" {
      newProps[v.Name] = newBool()
    } else if v.Type == "email" {
      newProps[v.Name] = newEmail(v.Length)
    }
  }

  return newProps
}

func GetObjs(data types.DataOpts) []map[string]interface{} {
  var newObjs []map[string]interface{}
  pop := data.Pop

  for i := 0;i < pop;i++ {
    newObjs = append(newObjs, GetSingleObj(data.Props))
  }

  return newObjs
}
