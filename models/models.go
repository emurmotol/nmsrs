package models

import (
	"errors"

	"github.com/zneyrl/nmsrs/helpers/lang"
)

var ErrInvalidObjectID = errors.New(lang.En["object_id_invalid"])
