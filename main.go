package main

import (
	"base-project/helper"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	helper.Panic(err)
}

func main() {

}
