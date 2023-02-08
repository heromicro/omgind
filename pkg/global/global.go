package global

import "C"
import (
	"encoding/json"
	"os"
	"strings"
	"sync"

	"github.com/go-redis/redis"
	"github.com/heromicro/omgind/pkg/config"
	"github.com/koding/multiconfig"
)

var (
	// CFG 全局配置(需要先执行MustLoad，否则拿不到配置)
	CFG  = new(config.Config)
	once sync.Once

	RdsCli redis.Cmdable
)

// PrintWithJSON 基于JSON格式输出配置
func PrintWithJSON() {
	if CFG.System.PrintConfig {
		b, err := json.MarshalIndent(CFG, "", " ")
		if err != nil {
			os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
			return
		}
		os.Stdout.WriteString(string(b) + "\n")
	}
}

// MustLoad 加载配置
func MustLoad(fpaths ...string) {
	once.Do(func() {
		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		}

		for _, fpath := range fpaths {
			if strings.HasSuffix(fpath, "toml") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "json") {
				loaders = append(loaders, &multiconfig.JSONLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "yaml") {
				loaders = append(loaders, &multiconfig.YAMLLoader{Path: fpath})
			}
		}

		m := multiconfig.DefaultLoader{
			Loader:    multiconfig.MultiLoader(loaders...),
			Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
		}
		m.MustLoad(CFG)
	})
}
