package fixtures

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/the-SkillConnect/SkillConnect/db"
)

var (
	Categories = [...]string{
		"tech",
		"database",
		"S3",
		"GraphQL",
		"GRPC",
	}
)

func AddUserIdentity(store db.Queries, ctx context.Context, i int) error {
	arg := db.InsertUserIdentityParams{
		Email:         fmt.Sprintf("user%d@james.com", i),
		Password:      fmt.Sprintf("abcd%d%d%d%d", i, i, i, i),
		Firstname:     fmt.Sprintf("user%d", i),
		Surname:       fmt.Sprintf("james%d", i),
		MobilePhone:   fmt.Sprintf("%d%d%d%d%d%d%d", i, i, i, i+6, i*2, i+3, i+2),
		WalletAddress: fmt.Sprintf("secureWallet%d%d%d%d%d%d%d", i, i, i, i+6, i*2, i+3, i+2),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	_, err := store.InsertUserIdentity(ctx, arg)
	return err
}

func AddUserProfile(store db.Queries, ctx context.Context, i int) error {
	arg := db.InsertUserProfileParams{
		UserID:        int64(i),
		Rating:        int64(rand.Intn(5)),
		Description:   sql.NullString{String: fmt.Sprintf("this is a mock description for user %d", i), Valid: true},
		DoneProjects:  int64(rand.Intn(20)),
		GivenProjects: int64(rand.Intn(100)),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	_, err := store.InsertUserProfile(ctx, arg)
	return err
}

func AddUserRecommendation(store db.Queries, ctx context.Context, i int) error {
	rID := rand.Intn(i) + 1
	arg := db.InsertUserRecommendationParams{
		GivenID:     int64(i),
		ReceivedID:  int64(rID),
		Description: fmt.Sprintf("this is recom from user%d to user%d", int64(i), int64(rID)),
	}
	_, err := store.InsertUserRecommendation(ctx, arg)
	return err
}

func AddProject(store db.Queries, ctx context.Context, i int) error {
	rID := rand.Intn(i) + 1
	arg := db.InsertProjectParams{
		Description: fmt.Sprintf("this is a mock mock mock mock mock project from user%d", int64(i)),
		Title:       fmt.Sprintf("this is a mock project title from user%d", int64(i)),
		TotalAmount: fmt.Sprintf("%.3f", rand.Float64()*float64(i*10)),
		DoneStatus:  sql.NullBool{Bool: false, Valid: true},
		UserID:      int64(rID),
		Fee:         fmt.Sprintf("%.3f", rand.Float64()*float64(i)),
		CategoryID:  sql.NullInt64{Int64: int64(rand.Intn(len(Categories) + 1)), Valid: true},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, err := store.InsertProject(ctx, arg)
	return err
}

func AddComment(store db.Queries, ctx context.Context, i int) error {
	rID := rand.Intn(i) + 1
	arg := db.InsertCommentParams{
		UserID:    int64(i),
		ProjectID: int64(rID),
		Date:      time.Now(),
		Text:      fmt.Sprintf("this a random %d comment", i),
	}
	_, err := store.InsertComment(ctx, arg)
	return err
}

func AddAssignProject(store db.Queries, ctx context.Context, i int) error {
	rID := rand.Intn(i) + 1

	arg := db.InsertAssignProjectParams{
		UserID:    int64(i),
		ProjectID: int64(rID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := store.InsertAssignProject(ctx, arg)
	return err
}

func AddCategory(store db.Queries, ctx context.Context) error {
	for _, value := range Categories {
		_, err := store.InsertCategory(ctx, value)
		if err != nil {
			return err
		}
	}
	return nil
}
