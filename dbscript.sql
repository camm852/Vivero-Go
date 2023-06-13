create database vivero;

\c vivero;

CREATE TABLE plants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    degree_survival FLOAT,
    amount_water_required FLOAT,
    amount_water_system FLOAT,
    degree_hydration INT,
    amount_nutrients_required INT,
    amount_nutrients_system FLOAT,
    degree_nutrition INT,
    created TIMESTAMPTZ DEFAULT current_timestamp,
    last_update TIMESTAMPTZ DEFAULT current_timestamp
);