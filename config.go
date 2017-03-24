package main

import(
  "fmt"
  "github.com/spf13/viper"
)

type Development struct{
  Port int
}

type AppConfig struct{
    Dev Development
}

func GetConfig() (*AppConfig, error){
  viper.SetConfigName("app_config")
  viper.AddConfigPath(".")
  if err := viper.ReadInConfig(); err != nil {
    return nil, err
  }
  var ac *AppConfig
  if err := viper.Unmarshal(&ac); err != nil {
    return nil, err
  }
  return ac, nil
}

func main(){
  ac, err := GetConfig()
  fmt.Println(err)
  fmt.Println(ac)
}
