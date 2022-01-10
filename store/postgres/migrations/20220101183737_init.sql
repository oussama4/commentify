-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(64) PRIMARY KEY,
	name VARCHAR(32) NOT NULL,
	email VARCHAR(64),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS pages (
	id VARCHAR(64) PRIMARY KEY,
    page_id VARCHAR(255),
	url TEXT NOT NULL,
	title TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS comments (
    id VARCHAR(64) PRIMARY KEY,
    body TEXT NOT NULL,
    parent_id VARCHAR(64),
    user_id VARCHAR(64) NOT NULL,
	page_id VARCHAR(64) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (page_id) REFERENCES pages (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS pages;
-- +goose StatementEnd
