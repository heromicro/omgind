package asyncq

import "github.com/google/wire"

// var ProviderSet = wire.NewSet(NewAsynqClientOpt, NewAsynqClient, NewAsynqServer)
var ProviderSet = wire.NewSet(NewAsynqClient)
