package main

import (
  "log"
  "todo-app"
  "todo-app/pkg/handler"
  "github.com/spf13/viper"
)

func main(){
  if err := initConfig(); err != nil{
    log.Fatalf("error initializing configs: %s", err.Error())
  }
  repos := repository.Repository()
  services := service.NewService(repos)
  handlers := handler.NewHandler(services)

  handlers := new(handler.Handler)
  srv := new(todo.Server)
  if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil{
    log.Fatalf("failed to initialize db: %s", err.Error())
  }
}

func initConfig() error {
  viper.AddConfigPath("configs")
  viper.SetConfigName("config")
  return viper.ReadInConfig()
}
