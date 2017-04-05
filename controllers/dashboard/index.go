package dashboard

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/shared/res"
)

func Index(w http.ResponseWriter, r *http.Request) {
	res.JSON(res.Make{http.StatusOK, "Gained access to protected resource", ""}, w)
}
