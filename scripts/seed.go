package main

import (
	"context"
	"log"
	"time"

	"github.com/the-SkillConnect/SkillConnect/db"
	"github.com/the-SkillConnect/SkillConnect/db/fixtures"
)

func main() {
	now := time.Now()
	var i int
	go func() {
		for {
			<-time.After(time.Millisecond * 500)
			log.Printf("Seed Elapsed: %v, inserted: %d\n", time.Since(now), i)
		}
	}()
	dbInstance, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	db.TearDown(dbInstance)
	db.CreateTables(dbInstance)
	ctx := context.Background()
	instance := db.New(dbInstance)
	fixtures.AddCategory(*instance, ctx)

	for i = 1; i < 10; i++ {
		fixtures.AddUserIdentity(*instance, ctx, i)
		fixtures.AddUserProfile(*instance, ctx, i)
		fixtures.AddUserRecommendation(*instance, ctx, i)
		fixtures.AddProject(*instance, ctx, i)
		fixtures.AddComment(*instance, ctx, i)
		fixtures.AddAssignProject(*instance, ctx, i)
	}
	log.Printf("Seed Ended in %v\n", time.Since(now))
}
