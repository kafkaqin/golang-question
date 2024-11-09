package main

import (
	"fmt"
	"golang-question/config"
	"golang-question/errorx"
)

const filePath = "demo.json"

type Secret string

func (s Secret) Validate() errorx.Error {
	if len(s) < 8 {
		return errorx.Cf(errorx.CODE_SECRET_TOO_SHORT, "invalid secret %s", s)
	}
	return nil
}

type Config struct {
	Secret Secret `yaml:"secret" json:"secret"`
}

var conf = config.Local[Config](filePath).Watch().InitData(Config{
	Secret: "hello world",
})

func main() {
	s := conf.Get().Secret
	if err := s.Validate(); err != nil {
		fmt.Printf("validate error: %+v\n", err)
	}
	fmt.Printf("update before: %+v\n", conf.Get())
	if err := conf.Update(Config{Secret: Secret("updated secret")}); err != nil {
		fmt.Printf("update error: %+v\n", err)
	}
	fmt.Printf("update after: %+v\n", conf.Get())

	if err := conf.Update(Config{Secret: Secret("s")}); err != nil {
		fmt.Printf("update error: %+v\n", err)
	}
	fmt.Printf("update after: %+v\n", conf.Get())

	s = conf.Get().Secret
	if err := s.Validate(); err != nil {
		fmt.Printf("validate error: %+v\n", err)
	}
}
