package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/the-SkillConnect/SkillConnect/db"
	"github.com/the-SkillConnect/SkillConnect/db/fixtures"
)

func main() {
	var insertedCount int = 5
	now := time.Now()

	go logSeedingProgress(now, &insertedCount)

	dbInstance, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	setupDatabase(dbInstance)

	instance := db.New(dbInstance)
	ctx := context.Background()

	seedDatabase(instance, ctx, &insertedCount)
	log.Printf("Seed Ended in %v\n", time.Since(now))
}

func setupDatabase(dbInstance *sql.DB) error {
	db.TearDown(dbInstance)
	return db.CreateTables(dbInstance)
}

func logSeedingProgress(startTime time.Time, insertedCount *int) {
	for {
		time.Sleep(500 * time.Millisecond)
		log.Printf("Seed Elapsed: %v, inserted: %d\n", time.Since(startTime), *insertedCount)
	}
}

func seedDatabase(instance *db.Queries, ctx context.Context, insertedCount *int) error {
	if err := fixtures.AddCategory(*instance, ctx); err != nil {
		return err
	}

	for i := 1; i < 1000; i++ {
		if err := addFixtureData(instance, ctx, i); err != nil {
			return err
		}
		*insertedCount = i
	}
	return nil
}

func addFixtureData(instance *db.Queries, ctx context.Context, i int) error {
	if err := fixtures.AddUserIdentity(*instance, ctx, i); err != nil {
		return err
	}
	if err := fixtures.AddUserProfile(*instance, ctx, i); err != nil {
		return err
	}
	if err := fixtures.AddUserRecommendation(*instance, ctx, i); err != nil {
		return err
	}
	if err := fixtures.AddProject(*instance, ctx, i); err != nil {
		return err
	}
	if err := fixtures.AddComment(*instance, ctx, i); err != nil {
		return err
	}
	return fixtures.AddAssignProject(*instance, ctx, i)
}