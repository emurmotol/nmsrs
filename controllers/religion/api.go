package religion

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/models/religion"
)

func All(w http.ResponseWriter, r *http.Request) {
	religs, err := religion.Search(bson.M{
		"name": bson.RegEx{
			Pattern: r.URL.Query().Get("q"),
			Options: "i",
		},
	})

	if err != nil {
		panic(err)
	}
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"religions": religs,
		},
	})
}
