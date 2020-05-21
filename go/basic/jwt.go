package main

import "fmt"
import "encoding/base64"
import "encoding/json"

type JwtClaims struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

func main() {
	claim := JwtClaims{"JWT", "HS256"}
	jsonStr, err := json.Marshal(claim)
	if err != nil {
		return
	}

	endcoded := base64.StdEncoding.EncodeToString(jsonStr)
	fmt.Println(endcoded)
  fmt.Printf("%x", endcoded)
}
