package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"task_1/database"
	"task_1/handler"
	"task_1/repositories"
	"task_1/services"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main(){

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".","_"))

	if _,err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port: viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	addr := ":" + config.Port
	fmt.Println("Server running di", addr)

	// setup database
	db,err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// setup repository, service, dan handler
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "OK",
		"message": "API Running",
	})
})

http.HandleFunc("/categories", categoryHandler.HandleCategories)

http.HandleFunc("/categories/", categoryHandler.HandleCategoriesByID)

	

	err = http.ListenAndServe(addr,nil)

	if err != nil{
		fmt.Print(("gagal running server"))
}
}



