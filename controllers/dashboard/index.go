package dashboard

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/shared/response"
)

func Index(w http.ResponseWriter, r *http.Request) {
	response.JSON(response.Make{http.StatusOK, "Gained access to protected resource", ""}, w)
}
