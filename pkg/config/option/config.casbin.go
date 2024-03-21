package option

// Casbin casbin配置参数
type CasbinConfig struct {
	Enable              bool
	Debug               bool
	ModelFile           string `default:"rbac_model.conf"`
	GenPolicyFile       string `default:"gen_rbac_policy.csv"`
	AutoLoad            bool
	AutoLoadInternal    int
	SkippedPathPrefixes []string
}
