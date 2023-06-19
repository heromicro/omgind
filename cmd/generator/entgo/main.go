package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {

	err := entc.Generate("./internal/scheme/entity", &gen.Config{
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
		Package: "github.com/heromicro/omgind/internal/gen/ent",
		Target:  "./internal/gen/ent",
		Templates: []*gen.Template{
			gen.MustParse(gen.NewTemplate("mutation_input").ParseDir("./internal/scheme/template")),
			gen.MustParse(gen.NewTemplate("stringer").ParseDir("./internal/scheme/template")),
			gen.MustParse(gen.NewTemplate("debug").ParseDir("./internal/scheme/template")),
		},
	},
	//  entc.Extensions(entviz.Extension{})
	)

	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
