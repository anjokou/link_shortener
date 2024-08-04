package main

import (
	"database/sql"
	"link_shortener/actions"
	"link_shortener/analytics"
	syserrors "link_shortener/errors"
	"link_shortener/links"
	"link_shortener/routers"
	"link_shortener/settings"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//The analytics service is currently in-memory so this keeps the services
	// in sync.
	if settings.ClearDatabaseOnStartup() {
		os.Remove(settings.SqliteFile())
	}

	db, err := sql.Open("sqlite3", settings.SqliteFile())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		panic(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://"+settings.MigrationsDir(),
		"sqlite3",
		driver,
	)
	if err != nil {
		panic(err)
	}

	err = migration.Up()

	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	hasher := links.CreateSha256Hasher()
	errorHandler := syserrors.CreateDefaultErrorHandler()

	linkDA := links.CreateSQLiteLinkDa(db, hasher, errorHandler)
	analytics := analytics.CreateRemoteAnalytics(errorHandler)
	defer analytics.Close()

	actions.InitLinkctions(linkDA, analytics)

	routers.InitLinks(hasher)

	server := gin.Default()
	routers.ApplyLinkRoutes(server)

	server.Run(":" + strconv.Itoa(settings.ListeningPort()))
}
