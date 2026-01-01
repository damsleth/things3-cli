package cli

import "os"

func authTokenFromEnv() string {
	return os.Getenv("THINGS_AUTH_TOKEN")
}
