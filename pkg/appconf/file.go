package appconf

import "github.com/sdual/mlserving/pkg/env"

type ConfFileType struct {
	extension string
}

func (cf ConfFileType) fileName(env env.SystemEnv) string {
	return string(env) + "." + cf.extension
}

var tomlFileType = ConfFileType{
	extension: "toml",
}
