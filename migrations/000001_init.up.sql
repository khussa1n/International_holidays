CREATE TABLE users (
    id                serial primary key,
    chat_id           integer unique not null,
    username          varchar(255) not null,
    first_query_time  varchar(255),
    all_queries_count integer
);

CREATE TABLE dates (
   id                serial primary key,
   chat_id           integer references users(chat_id),
   description       text not null,
   date          varchar(30)
);