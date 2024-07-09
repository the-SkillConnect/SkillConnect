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

CREATE TABLE IF NOT EXISTS comment (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    project_id BIGINT NOT NULL,
    date TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    text TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS assign_project (
    user_id BIGINT NOT NULL,
    project_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, project_id)
);

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

CREATE TABLE IF NOT EXISTS user_profile (
    user_id BIGINT PRIMARY KEY,
    rating BIGINT NOT NULL DEFAULT 0,
    description TEXT NOT NULL,
    done_project BIGINT NOT NULL DEFAULT 0,
    given_project BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_recommendation (
    given_id BIGINT NOT NULL,
    received_id BIGINT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (given_id, received_id)
);

CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL UNIQUE
);

ALTER TABLE project ADD CONSTRAINT project_fk_user FOREIGN KEY (user_id) REFERENCES user_identity(id);
ALTER TABLE project ADD CONSTRAINT project_fk_category FOREIGN KEY (categories) REFERENCES category(id);

ALTER TABLE comment ADD CONSTRAINT comment_fk_user FOREIGN KEY (user_id) REFERENCES user_identity(id);
ALTER TABLE comment ADD CONSTRAINT comment_fk_project FOREIGN KEY (project_id) REFERENCES project(id);

ALTER TABLE assign_project ADD CONSTRAINT assign_project_fk_user FOREIGN KEY (user_id) REFERENCES user_identity(id);
ALTER TABLE assign_project ADD CONSTRAINT assign_project_fk_project FOREIGN KEY (project_id) REFERENCES project(id);

ALTER TABLE user_profile ADD CONSTRAINT user_profile_fk_user FOREIGN KEY (user_id) REFERENCES user_identity(id);
ALTER TABLE user_profile ADD CONSTRAINT user_profile_fk_recommendation FOREIGN KEY (recommendation_id) REFERENCES user_recommendation(given_id);

ALTER TABLE user_recommendation ADD CONSTRAINT user_recommendation_fk_given FOREIGN KEY (given_id) REFERENCES user_identity(id);
ALTER TABLE user_recommendation ADD CONSTRAINT user_recommendation_fk_received FOREIGN KEY (received_id) REFERENCES user_identity(id);
