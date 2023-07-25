package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	// "github.com/flume/enthistory"
)

func main() {

	err := entc.Generate("./internal/scheme/entity",
		&gen.Config{
			Features: []gen.Feature{
				//gen.FeatureEntQL,
				// gen.FeatureSnapshot,
				// gen.FeatureSchemaConfig,
				gen.FeatureLock,
				gen.FeatureModifier,
				gen.FeatureUpsert,
				gen.FeatureIntercept,
				// entc.FeatureNames("intercept"),
			},
			Package: "github.com/heromicro/omgind/internal/gen/mainent",
			Target:  "./internal/gen/mainent",
			Templates: []*gen.Template{
				gen.MustParse(gen.NewTemplate("mutation_input").ParseDir("./internal/scheme/template")),
				gen.MustParse(gen.NewTemplate("stringer").ParseDir("./internal/scheme/template")),
				gen.MustParse(gen.NewTemplate("debug").ParseDir("./internal/scheme/template")),
			},
		},

		entc.Extensions(
		// enthistory.NewHistoryExtension(
		// 	enthistory.WithUpdatedBy("userid", enthistory.ValueTypeString),
		// 	enthistory.WithAuditing(),
		// 	enthistory.WithHistoryTimeIndex(),
		// ),
		),

	//  entc.Extensions(entviz.Extension{})
	)

	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
