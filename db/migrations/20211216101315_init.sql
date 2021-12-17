-- migrate:up

CREATE TABLE products
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP(0) WITH TIME ZONE
);

CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    phone      VARCHAR(11) NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP(0) WITH TIME ZONE
);

CREATE INDEX idx_users__phone
    on users (phone);

CREATE UNIQUE INDEX unq_users__phone
    on users (phone)
    where (deleted_at IS NULL);

CREATE TABLE phone_auth_request
(
    id         SERIAL PRIMARY KEY,
    phone      VARCHAR(11) NOT NULL,
    code       VARCHAR(4)  NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE DEFAULT now(),
    expired_at TIMESTAMP(0) WITH TIME ZONE DEFAULT now() + interval '2 minute',
    ip         inet,
    success    bool        NOT NULL        DEFAULT FALSE
);

CREATE INDEX idx_phone_auth_requests__phone
    on phone_auth_request (phone);

CREATE TABLE phone_auth_request_attempts
(
    id                     SERIAL PRIMARY KEY,
    phone_auth_requests_id INT        NOT NULL,
    input_code             VARCHAR(4) NOT NULL,
    created_at             TIMESTAMP(0) WITH TIME ZONE DEFAULT now(),
    ip                     inet,
    success                bool       NOT NULL         DEFAULT FALSE
);

-- ALTER TABLE phone_auth_request_attempts ADD CONSTRAINT fk_phone_auth_request_attempts__phone_auth_requests_id
-- FOREIGN KEY (phone_auth_requests_id) REFERENCES phone_auth_request(id);

-- migrate:down

DROP TABLE IF EXISTS phone_auth_request_attempts;
DROP TABLE IF EXISTS phone_auth_request;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS products;

