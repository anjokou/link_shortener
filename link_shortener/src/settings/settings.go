package settings

import (
	"os"
	"strconv"
)

func ListeningPort() int {
	env, found := os.LookupEnv("HTTP_PORT")

	if found {
		parsedEnv, err := strconv.Atoi(env)
		if err == nil {
			return parsedEnv
		}
	}

	return 8000
}

func SqliteFile() string {
	env, found := os.LookupEnv("SQLITE_FILE")
	if found {
		return env
	} else {
		return "database.db"
	}
}

func MigrationsDir() string {
	env, found := os.LookupEnv("MIGRATIONS_DIR")
	if found {
		return env
	} else {
		return "../migrations"
	}
}

func AnalyticsHost() string {
	env, found := os.LookupEnv("ANALYTICS_HOST")
	if found {
		return env
	} else {
		return "localhost"
	}
}

func AnalyticsPort() int {
	env, found := os.LookupEnv("ANALYTICS_PORT")

	if found {
		parsedEnv, err := strconv.Atoi(env)
		if err == nil {
			return parsedEnv
		}
	}

	return 8002
}

func ClearDatabaseOnStartup() bool {
	_, found := os.LookupEnv("CLEAR_DB_ON_STARTUP")
	return found
}
