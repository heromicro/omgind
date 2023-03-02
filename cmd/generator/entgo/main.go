package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
)

func main() {

	err := entc.Generate("./internal/schema/entity", &gen.Config{
		Features: []gen.Feature{
			//gen.FeatureEntQL,
			//gen.FeatureSnapshot,
			// gen.FeatureSchemaConfig,
			// gen.FeatureLock,
			// gen.FeatureModifier,
			gen.FeatureUpsert,
		},
		Package: "github.com/heromicro/omgind/internal/gen/ent",
		Target:  "./internal/gen/ent",
		Templates: []*gen.Template{
			gen.MustParse(gen.NewTemplate("mutation_input").ParseDir("./internal/schema/template")),
			gen.MustParse(gen.NewTemplate("stringer").ParseDir("./internal/schema/template")),
		},
	}, entc.Extensions(entviz.Extension{}))

	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
