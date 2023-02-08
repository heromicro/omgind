package ulid

import (
	"errors"
	mathrand "math/rand"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/heromicro/omgind/pkg/helper/uid"
)

// ErrGeneratingID indicates error in generating ULID
var ErrGeneratingID = errors.New("generating id failed")

var _ uid.IDProvider = (*ulidProvider)(nil)

type ulidProvider struct {
	entropy *mathrand.Rand
}

func New() uid.IDProvider {
	seed := time.Now().UnixNano()
	source := mathrand.NewSource(seed)
	return &ulidProvider{
		entropy: mathrand.New(source),
	}
}

func (up *ulidProvider) ID() (string, error) {
	id, err := ulid.New(ulid.Timestamp(time.Now()), up.entropy)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// MustString Create ulid
func MustString() string {
	idp := New()
	id, err := idp.ID()
	if err != nil {
		panic(err)
	}
	return id
}
