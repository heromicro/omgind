package uuid

import (
	"errors"

	"github.com/google/uuid"

	"github.com/heromicro/omgind/pkg/helper/uid"
)

var ErrGeneratingID = errors.New("generating id failed")
var _ uid.IDProvider = (*uuidProvider)(nil)

type uuidProvider struct {
}

func New() uid.IDProvider {
	return &uuidProvider{}
}

func (up *uuidProvider) ID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// MustString Create uuid
func MustString() string {
	idp := New()
	id, err := idp.ID()
	if err != nil {
		panic(err)
	}
	return id
}
