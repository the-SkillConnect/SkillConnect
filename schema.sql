CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    firstname VARCHAR(100),
    surname VARCHAR(100),
    mobile_phone VARCHAR(20) UNIQUE,
    role_id INTEGER REFERENCES Role(id)
);

CREATE TABLE IF NOT EXISTS Role (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) UNIQUE NOT NULL
);

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

CREATE TABLE IF NOT EXISTS ProjectComment (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id) ON DELETE SET NULL,
    project_id INTEGER REFERENCES Project(id) ON DELETE CASCADE,
    date TIMESTAMP,
    text TEXT
);

CREATE TABLE IF NOT EXISTS AssignedProject (
    user_id INTEGER REFERENCES Users(id) ON DELETE CASCADE,
    project_id INTEGER REFERENCES Project(id) ON DELETE CASCADE,
    issued BOOLEAN,
    PRIMARY KEY (user_id, project_id)
);
