CREATE TABLE users
(
    tg_user_id BIGINT PRIMARY KEY,
    lang       VARCHAR(5) NOT NULL DEFAULT 'en'
);

CREATE TABLE playlist_info
(
    id             SERIAL PRIMARY KEY,
    playlist_title VARCHAR(255) NOT NULL,
    video_id       VARCHAR(255) NOT NULL,
    video_title    VARCHAR(255) NOT NULL,
    owner_id       VARCHAR(255) NOT NULL,
    owner_title    VARCHAR(255) NOT NULL,
    playlist_id    VARCHAR(255) NOT NULL
);
