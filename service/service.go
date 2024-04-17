package service

import (
	"math/rand"
)

func newStr() string {
  chars := [26]byte{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'}

  var str []byte

  for i := 0;i < 5;i++ {
    str = append(str, chars[rand.Intn(26)])
  }

  return string(str)
}
