package common

import (
	"github.com/heromicro/omgind/internal/gen/ent"
	"github.com/heromicro/omgind/internal/scheme"
	"github.com/heromicro/omgind/pkg/config"
)

func MakeEntClient(cf string) (*ent.Client, func(), error) {

	vip, err := config.New(cf)
	if err != nil {
		return nil, nil, err
	}
	appconfig := config.NewConfig(vip)

	entClient, cleanup, err := scheme.New(appconfig)

	if err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {
		cleanup()
	}

	return entClient, cleanFunc, nil
}
