package models

import (
	"errors"

	"github.com/emurmotol/nmsrs/helpers/lang"
)

var ErrInvalidObjectID = errors.New(lang.En["ObjectIdInvalid"])
