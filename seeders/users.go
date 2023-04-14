package seeders

import (
	"github.com/9ziggy9/go-starter/schema"
)

var Users []*schema.User = []*schema.User{
	&schema.User{Name: "ziggy", Password: "pass"},
	&schema.User{Name: "charlie", Password: "pass"},
	&schema.User{Name: "echo", Password: "pass"},
}
