package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	ErrorCreateUserIdentityTable  = errors.New("could not create user identity table")
	ErrorCreateUserProfileTable   = errors.New("could not create user profile table")
	ErrorCreateUserRecommendation = errors.New("could not create user recommendation table")
	ErrorCreateProjectTable       = errors.New("could not create project table")
	ErrorCreateComment            = errors.New("could not create comment table")
	ErrorCreateCategory           = errors.New("could not create category table")
	ErrorCreateAssignedProject    = errors.New("could not create assigned project table")
)

type DBParameter struct {
	Address  string
	Username string
	Password string
}

func NewDBParameter(param DBParameter) *DBParameter {
	return &DBParameter{
		Address:  param.Address,
		Username: param.Username,
		Password: param.Password,
	}
}

func InitDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}

	params := DBParameter{
		Address:  os.Getenv("DB_LISTEN_ADDRESS"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	connection := fmt.Sprintf("postgres://%s:%s@%s/database?sslmode=disable", params.Username, params.Password, params.Address)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(10)

	return db, err
}

func CreateTables(db *sql.DB) error {
	err := CreateUserIdentityTable(db)
	if err != nil {
		return err
	}

	err = CreateUserProfileTable(db)
	if err != nil {
		return err
	}

	err = CreateUserRecommendationTable(db)
	if err != nil {
		return err
	}

	err = CreateProjectTable(db)
	if err != nil {
		return err
	}

	err = CreateCommentTable(db)
	if err != nil {
		return err
	}

	err = CreateCategoryTable(db)
	if err != nil {
		return err
	}

	err = CreateAssignedProjectTable(db)
	return err
}

func TearDown(db *sql.DB) {
	db.Exec("DROP TABLE IF EXISTS user_identity")
	db.Exec("DROP TABLE IF EXISTS user_profile")
	db.Exec("DROP TABLE IF EXISTS user_recommendation")
	db.Exec("DROP TABLE IF EXISTS project")
	db.Exec("DROP TABLE IF EXISTS comment")
	db.Exec("DROP TABLE IF EXISTS assign_project")
	db.Exec("DROP TABLE IF EXISTS category")
}

func CreateUserIdentityTable(db *sql.DB) error {
	createUserIdentityTable := `
	CREATE TABLE IF NOT EXISTS user_identity (
		id BIGSERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		encrypted_password TEXT NOT NULL,
		first_name TEXT NOT NULL,
		surname TEXT NOT NULL,
		mobile_phone TEXT NOT NULL UNIQUE,
		wallet_address TEXT NOT NULL UNIQUE,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(createUserIdentityTable)
	if err != nil {
		return errors.Join(ErrorCreateUserIdentityTable, err)
	}
	return nil
}

func CreateUserProfileTable(db *sql.DB) error {
	createUserProfileTable := `
	CREATE TABLE IF NOT EXISTS user_profile (
		user_id BIGINT PRIMARY KEY,
		rating BIGINT NOT NULL DEFAULT 0,
		description TEXT NOT NULL,
		done_project BIGINT NOT NULL DEFAULT 0,
		given_project BIGINT NOT NULL DEFAULT 0,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(createUserProfileTable)
	if err != nil {
		return errors.Join(ErrorCreateUserProfileTable, err)
	}
	return nil
}

func CreateUserRecommendationTable(db *sql.DB) error {
	createUserRecommendationTable := `
	CREATE TABLE IF NOT EXISTS user_recommendation (
		given_id BIGINT NOT NULL,
		received_id BIGINT NOT NULL,
		description TEXT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (given_id, received_id)
	);
	`
	_, err := db.Exec(createUserRecommendationTable)
	if err != nil {
		return errors.Join(ErrorCreateUserRecommendation, err)
	}
	return nil
}

func CreateProjectTable(db *sql.DB) error {
	createProjectTable := `
	CREATE TABLE IF NOT EXISTS project (
		id BIGSERIAL PRIMARY KEY,
		description TEXT NOT NULL,
		title TEXT NOT NULL,
		total_amount NUMERIC(10,2) NOT NULL,
		done_status BOOLEAN DEFAULT FALSE,
		user_id BIGINT NOT NULL,
		fee NUMERIC(10,2) NOT NULL,
		category_id BIGINT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(createProjectTable)
	if err != nil {
		return errors.Join(ErrorCreateProjectTable, err)
	}
	return nil
}

func CreateCommentTable(db *sql.DB) error {
	createComment := `
	CREATE TABLE IF NOT EXISTS comment (
		id BIGSERIAL PRIMARY KEY,
		user_id BIGINT NOT NULL,
		project_id BIGINT NOT NULL,
		date TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		text TEXT NOT NULL
	);	
	`
	_, err := db.Exec(createComment)
	if err != nil {
		return errors.Join(ErrorCreateComment, err)
	}
	return nil
}

func CreateCategoryTable(db *sql.DB) error {
	createCategory := `
	CREATE TABLE IF NOT EXISTS category (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL UNIQUE
	);
	`
	_, err := db.Exec(createCategory)
	if err != nil {
		return errors.Join(ErrorCreateCategory, err)
	}
	return nil
}

func CreateAssignedProjectTable(db *sql.DB) error {
	createAssignedProject := `
	CREATE TABLE IF NOT EXISTS assign_project (
		user_id BIGINT NOT NULL,
		project_id BIGINT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (user_id, project_id)
	);
	`
	_, err := db.Exec(createAssignedProject)
	if err != nil {
		return errors.Join(ErrorCreateAssignedProject, err)
	}
	return err
}
