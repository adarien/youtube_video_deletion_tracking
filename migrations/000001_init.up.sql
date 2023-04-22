CREATE TABLE users
(
    tg_user_id BIGINT PRIMARY KEY,
    lang       VARCHAR(5) NOT NULL DEFAULT 'en'
);
