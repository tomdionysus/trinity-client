package parser

import (
	// "errors"
	"github.com/tomdionysus/trinity/sql"
	"regexp"
)

const (
	CREATE = iota
	SELECT = iota
	INSERT = iota
	UPDATE = iota
	DELETE = iota
	ALTER  = iota
)

func Parse(input *string) (sql.Term, error) {

	// Get Command
	_ = regexp.MustCompile("([A-Za-Z]+).*")

	return nil, nil
}
