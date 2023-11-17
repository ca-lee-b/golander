CREATE TABLE events(
    id SERIAL ,

    title TEXT NOT NULL,
    owner_id UUID NOT NULL,

    created_at DATE NOT NULL DEFAULT CURRENT_DATE,
    available_dates DATE[],

    participants UUID[],

    CONSTRAINT fk_owner FOREIGN KEY(owner_id) REFERENCES users(id),
    PRIMARY KEY (id)
);