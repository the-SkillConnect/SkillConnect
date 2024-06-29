package db

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

var (
	ErrorCreateUserTable       = errors.New("could not create user table")
	ErrorRoleTable             = errors.New("could not create Role table")
	ErrorCreateProjectTable    = errors.New("could not project table")
	ErrorCreateProjectComment  = errors.New("could not project comment table")
	ErrorCreateAssignedProject = errors.New("could not project comment table")
)

func InitDB() (*sql.DB, error) {
	// Establish database connection
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost:5432/database?sslmode=disable")
	if err != nil {
		return nil, err
	}

	// Ping the database to verify connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(10)

	err = CreateRoleTable(db)
	if err != nil {
		return nil, err
	}

	err = CreateUserTable(db)
	if err != nil {
		return nil, err
	}

	err = CreateProjectTable(db)
	if err != nil {
		return nil, err
	}

	err = CreateProjectComment(db)
	if err != nil {
		return nil, err
	}
	err = CreateAssignedProject(db)

	return db, err
}

func TearDown(db *sql.DB) {
	db.Exec("DROP TABLE IF EXISTS Role")
	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("DROP TABLE IF EXISTS Project")
	db.Exec("DROP TABLE IF EXISTS ProjectComment")
	db.Exec("DROP TABLE IF EXISTS AssignedProject")

}

func CreateUserTable(db *sql.DB) error {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		firstname VARCHAR(100),
		surname VARCHAR(100),
		mobile_phone VARCHAR(20) UNIQUE,
		role_id INTEGER REFERENCES Role(id)
	);
	`
	_, err := db.Exec(createUserTable)
	if err != nil {
		return errors.Join(ErrorCreateUserTable, err)
	}
	return nil
}

func CreateRoleTable(db *sql.DB) error {
	createRoleTable := `
	CREATE TABLE IF NOT EXISTS Role (
		id SERIAL PRIMARY KEY,
		type VARCHAR(50) UNIQUE NOT NULL
	);
	`
	_, err := db.Exec(createRoleTable)
	if err != nil {
		return errors.Join(ErrorRoleTable, err)
	}
	return nil
}

func CreateProjectTable(db *sql.DB) error {
	createProjectTable := `
	CREATE TABLE IF NOT EXISTS Project (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255),
		description TEXT,
		total_amount INTEGER, 
		order_date DATE,
		status BOOLEAN,
		user_id INTEGER REFERENCES Users(id) ON DELETE CASCADE,
		fee INTEGER
	);
	`
	_, err := db.Exec(createProjectTable)
	if err != nil {
		return errors.Join(ErrorCreateProjectTable, err)
	}
	return nil
}

func CreateProjectComment(db *sql.DB) error {
	createProjectComment := `
	CREATE TABLE IF NOT EXISTS ProjectComment (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES Users(id) ON DELETE SET NULL,
		project_id INTEGER REFERENCES Project(id) ON DELETE CASCADE,
		date TIMESTAMP,
		text TEXT
	);

	`
	_, err := db.Exec(createProjectComment)
	if err != nil {
		return errors.Join(ErrorCreateProjectComment, err)
	}
	return nil
}

func CreateAssignedProject(db *sql.DB) error {
	createAssignedProject := `
	CREATE TABLE IF NOT EXISTS AssignedProject (
		user_id INTEGER REFERENCES Users(id) ON DELETE CASCADE,
		project_id INTEGER REFERENCES Project(id) ON DELETE CASCADE,
		issued BOOLEAN,
		PRIMARY KEY (user_id, project_id)
	);
	`
	_, err := db.Exec(createAssignedProject)
	if err != nil {
		return errors.Join(ErrorCreateAssignedProject, err)
	}
	return err
}
