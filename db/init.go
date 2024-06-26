package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
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
		log.Fatal(err)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(10)

	CreateRoleTable(db)
	CreateUserTable(db)
	CreateProjectTable(db)
	CreateProjectComment(db)
	CreateAssignedProject(db)

	return db, nil
}

func TearDown(db *sql.DB) {
	db.Exec("DROP TABLE IF EXISTS Role")
	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("DROP TABLE IF EXISTS Project")
	db.Exec("DROP TABLE IF EXISTS ProjectComment")
	db.Exec("DROP TABLE IF EXISTS AssignedProject")

}

func CreateUserTable(db *sql.DB) {
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
		log.Fatal("Could not create user  table " + err.Error())
	}
}

func CreateRoleTable(db *sql.DB) {
	createRoleTable := `
	CREATE TABLE IF NOT EXISTS Role (
		id SERIAL PRIMARY KEY,
		type VARCHAR(50) UNIQUE NOT NULL
	);
	`
	_, err := db.Exec(createRoleTable)
	if err != nil {
		log.Fatal("Could not create Role table " + err.Error())
	}
}

func CreateProjectTable(db *sql.DB) {
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
		log.Fatal("Could not project table " + err.Error())
	}
}

func CreateProjectComment(db *sql.DB) {
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
		log.Fatal("Could not project comment table " + err.Error())
	}
}

func CreateAssignedProject(db *sql.DB) {
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
		log.Fatal("Could not project comment table " + err.Error())
	}
}
