CREATE TABLE IF NOT EXISTS expenses (
    id BIGSERIAL PRIMARY KEY ,
    uuid TEXT NOT NULL ,
    publication_date TEXT NOT NULL,
    fiscal_year TEXT NOT NULL,
    agency_code TEXT NOT NULL,
    agency_name TEXT NOT NULL,
    total_fund TEXT NOT NULL,
    city_fund TEXT NOT NULL,
    remark TEXT
);