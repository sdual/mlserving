package env

import "os"

type SystemEnv string

const (
	Prd  SystemEnv = "prd"
	Stg  SystemEnv = "stg"
	Dev  SystemEnv = "dev"
	Test SystemEnv = "test"

	// envVal is the variable name which represents the current environment, prd, stg, and so on.
	envVal = "SYSTEM_ENV"
)

func CurrentEnv() SystemEnv {
	envName := os.Getenv(envVal)
	switch envName {
	case "prd":
		return Prd
	case "stg":
		return Stg
	case "dev":
		return Dev
	default:
		return Test
	}
}
