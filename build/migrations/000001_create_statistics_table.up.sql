CREATE TABLE IF NOT EXISTS statisctics(
    id SERIAL,
    date date NOT NULL PRIMARY KEY,
    views int DEFAULT 0,
    clicks int DEFAULT 0,
    cost numeric DEFAULT 0
);

CREATE INDEX statisctics_date_index ON statisctics (date);