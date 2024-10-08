package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	"github.com/tvgelderen/fiscora/auth"
	"github.com/tvgelderen/fiscora/config"
	"github.com/tvgelderen/fiscora/handlers"
	"github.com/tvgelderen/fiscora/seed"
)

func main() {
	env := config.Envs
	if env.DBConnectionString == "" {
		log.Fatal("No database connection string found")
	}

	connection, err := sql.Open("postgres", env.DBConnectionString)
	if err != nil {
		log.Fatalf("Error establishing database connection: %s", err.Error())
	}

	seedFlag := flag.Bool("seed", false, "Set to true to seed the demo account")
	seedMyAccountFlag := flag.Bool("seed-me", false, "Set to true to seed my account too")
	flag.Parse()
	if *seedFlag {
		seed.Seed(connection)
	}
	if *seedMyAccountFlag {
		seed.SeedMyAccount(connection)
	}

	authService := auth.NewAuthService()
	handler := handlers.NewAPIHandler(connection, authService)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	base := e.Group("/api")

	// Enable users to view the site without logging in
	base.GET("/auth/demo", handler.HandleDemoLogin)

	base.GET("/auth/:provider", handler.HandleOAuthLogin)
	base.GET("/auth/callback/:provider", handler.HandleOAuthCallback)
	base.GET("/auth/logout", handler.HandleLogout, handler.AuthorizeEndpoint)
	base.GET("/transactions/types/intervals", handler.HandleGetTransactionIntervals)
	base.GET("/transactions/types/income", handler.HandleGetIncomeTypes)
	base.GET("/transactions/types/expense", handler.HandleGetExpenseTypes)

	users := base.Group("/users", handler.AuthorizeEndpoint)
	users.GET("/me", handler.HandleGetMe)

	transactions := base.Group("/transactions", handler.AuthorizeEndpoint)
	transactions.GET("", handler.HandleGetTransactions)
	transactions.POST("", handler.HandleCreateTransaction)
	transactions.PUT("/:id", handler.HandleUpdateTransaction)
	transactions.DELETE("/:id", handler.HandleDeleteTransaction)
	transactions.GET("/summary/month", handler.HandleGetTransactionMonthInfo)
	transactions.GET("/summary/month/type", handler.HandleGetTransactionsPerType)
	transactions.GET("/summary/year", handler.HandleGetTransactionYearInfo)
	transactions.GET("/summary/year/type", handler.HandleGetTransactionsYearInfoPerType)

	e.Logger.Fatal(e.Start(env.Port))
}
