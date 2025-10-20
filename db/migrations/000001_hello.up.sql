CREATE TABLE IF NOT EXISTS hello(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    word VARCHAR(128) NOT NULL
);

INSERT INTO hello (word) VALUES ('world');
