CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    avatar VARCHAR(255),
    first_name VARCHAR(100),
    last_name  VARCHAR(100),
    middle_name VARCHAR(100),
    email VARCHAR(100) NOT NULL,
    email_verified_at VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    is_banned BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
);

CREATE UNIQUE INDEX unique_user_email 
ON TABLE users(email); 


INSERT INTO public.users
(avatar, first_name, last_name, middle_name, email, email_verified_at, password, is_banned, created_at, updated_at, deleted_at)
VALUES('https://i.pravatar.cc/150?img=66', 'firstname', 'lastname', 'middle', 'cheeeasy2501@gmail.com', '', '', false, '', '', '');
