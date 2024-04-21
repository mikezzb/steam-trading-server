package app

import (
	"log"

	"github.com/beego/beego/v2/core/validation"
)

func MakeErrors(err []*validation.Error) {
	for _, e := range err {
		log.Println(e.Key, e.Message)
	}
}
