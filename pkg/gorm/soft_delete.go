package gorm

import "github.com/golang-module/carbon/v2"

// DeletedAt auto sign item soft deleted time
// !TODO finish gorm deleted at time implement
type DeletedAt struct {
	carbon.DateTime
}
