package utils

import "os"

func LookupOrDefaultEnv(defaultVal string, envKey string) string {
    envVal, exists := os.LookupEnv(envKey)
    if exists {
        return envVal
    } else {
        return defaultVal
    }
}