-- postgres database

CREATE TABLE roots(
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    wazan INTEGER NOT NULL,
    word TEXT NOT NULL,
    masdar TEXT NOT NULL
);

CREATE UNIQUE INDEX ON roots(word);