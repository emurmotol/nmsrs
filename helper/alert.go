package helper

import (
	"fmt"
)

type Alert struct {
	Type    string
	Content string
}

func (alert *Alert) String() string {
	markup := fmt.Sprintf(`<div class="alert alert-%s alert-dismissible" role="alert">
	<button type="button" class="close" data-dismiss="alert" aria-label="Close">
	<span aria-hidden="true">&times;</span></button>%s</div>`, alert.Type, alert.Content)
	return markup
}
