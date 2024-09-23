CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);

INSERT INTO users (id, name, email) VALUES ('1', 'John Doe', 'test@test.com');
