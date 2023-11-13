CREATE TABLE Users (
    id uuid DEFAULT uuid_generate_v4(),
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,

    PRIMARY KEY (id)
);