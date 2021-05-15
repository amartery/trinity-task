CREATE TABLE Users
(
    id   serial not null unique,
    name varchar(24) not null
);

CREATE TABLE Comments
(
    id   serial not null unique,
    video_id int not null,
    text varchar(255),
    user_id int references Users (id) on delete cascade not null,
    create_at timestamp without time zone not null
);

CREATE TABLE Likes
(
    video_id int not null,
    user_id int references Users (id) on delete cascade not null,
    create_at timestamp without time zone not null
);