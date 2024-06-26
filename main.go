package main

import (
	"log"

	"github.com/the-SkillConnect/SkillConnect/db"
)

func main() {
	_, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

}
