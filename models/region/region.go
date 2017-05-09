package region

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Region struct {
	ObjectID       bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Code     string        `schema:"code" json:"code" bson:"code,omitempty"`
	Desc     string        `schema:"desc" json:"desc" bson:"desc,omitempty"`
	PSGCCode string        `schema:"psgc_code" json:"psgc_code" bson:"psgcCode,omitempty"`
}

func All() ([]Region, error) {
	regs := []Region{}

	if err := db.Regions.Find(nil).Sort("+desc").All(&regs); err != nil {
		return nil, err
	}
	return regs, nil
}

func (reg *Region) Insert() (string, error) {
	reg.ObjectID = bson.NewObjectId()

	if err := db.Regions.Insert(reg); err != nil {
		return "", err
	}
	return reg.ObjectID.Hex(), nil
}

func FindByID(id string) (*Region, error) {
	var reg Region

	if !bson.IsObjectIdHex(id) {
		return &reg, models.ErrInvalidObjectID
	}

	if err := db.Regions.FindId(bson.ObjectIdHex(id)).One(&reg); err != nil {
		return &reg, err
	}
	return &reg, nil
}

func FindAllBy(key string, value interface{}) ([]Region, error) {
	regs := []Region{}

	if err := db.Regions.Find(bson.M{key: value}).Sort("+desc").All(&regs); err != nil {
		return regs, err
	}
	return regs, nil
}
