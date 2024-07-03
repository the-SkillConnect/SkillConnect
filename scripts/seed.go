package main

import (
	"context"
	"fmt"
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
			fmt.Printf("Seed Elapsed: %v, inserted: %d\n", time.Since(now), i)
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

<<<<<<< Updated upstream
	for i := 1; i < 1000; i++ {
		if err := addFixtureData(instance, ctx, i); err != nil {
			return err
=======
	for i = 1; i < 10; i++ {
		fixtures.AddUserIdentity(*instance, ctx, i)
		err := fixtures.AddUserProfile(*instance, ctx, i)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
>>>>>>> Stashed changes
		}
		fixtures.AddUserRecommendation(*instance, ctx, i)
		fixtures.AddProject(*instance, ctx, i)
		fixtures.AddComment(*instance, ctx, i)
		fixtures.AddAssignProject(*instance, ctx, i)
	}
<<<<<<< Updated upstream
	return fixtures.AddAssignProject(*instance, ctx, i)
}
=======
	fmt.Printf("Seed Ended in %v\n", time.Since(now))
}
>>>>>>> Stashed changes
