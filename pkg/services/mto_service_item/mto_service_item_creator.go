package mtoserviceitem

import (
	"fmt"

	"github.com/gobuffalo/validate"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
)

type createMTOServiceItemQueryBuilder interface {
	CreateOne(model interface{}) (*validate.Errors, error)
}

type mtoServiceItemCreator struct {
	builder createMTOServiceItemQueryBuilder
}

func (o *mtoServiceItemCreator) CreateMTOServiceItem(serviceItem *models.MtoServiceItem) (*models.MtoServiceItem, *validate.Errors, error) {
	verrs, err := o.builder.CreateOne(serviceItem)
	fmt.Println(verrs, err)
	if verrs != nil || err != nil {
		return nil, verrs, err
	}

	return serviceItem, nil, nil
}

func NewMTOServiceItemCreator(builder createMTOServiceItemQueryBuilder) services.MTOServiceItemCreator {
	return &mtoServiceItemCreator{builder}
}
