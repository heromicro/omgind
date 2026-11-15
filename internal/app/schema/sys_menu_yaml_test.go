package schema

import (
	"gopkg.in/yaml.v2"
	"os"
	"testing"
)

func TestMenuTreesDecodeConfig(t *testing.T) {
	for _, name := range []string{"menu.yaml", "menu.origin.yaml"} {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile("../../../configs/" + name)
			if err != nil {
				t.Fatal(err)
			}
			var menus MenuTrees
			if err := yaml.UnmarshalStrict(data, &menus); err != nil {
				t.Fatal(err)
			}
			if len(menus) == 0 {
				t.Fatal("expected menus")
			}
		})
	}
}
