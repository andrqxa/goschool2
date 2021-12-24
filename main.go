package main

import (
	"fmt"
	"net/http"

	"github.com/andrqxa/goschool2/controllers"
	"github.com/andrqxa/goschool2/models"
	"github.com/gorilla/mux"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// /*
func main() {
	us, err := models.NewUserService()
	must(err)

	_ = controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.HandleFunc("/", usersC.New).Methods("GET")
	r.HandleFunc("/", usersC.Create).Methods("POST")
	fmt.Println("Starting the server on :3000...")

	http.ListenAndServe(":3000", r)
}

// */
/*
var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
*/
