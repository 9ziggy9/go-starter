package seeders

import (
	"github.com/9ziggy9/go-starter/schema"
	"time"
)

var Users []*schema.User = []*schema.User{
	&schema.User{
		Name: "ziggy", Password: "pass",
		CreatedAt: time.Now(), UpdatedAt: time.Now(),
	},
	&schema.User{
		Name: "charlie", Password: "pass",
		CreatedAt: time.Now(), UpdatedAt: time.Now(),
	},
	&schema.User{
		Name: "echo", Password: "pass",
		CreatedAt: time.Now(), UpdatedAt: time.Now(),
	},
}
