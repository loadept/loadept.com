CREATE TABLE IF NOT EXISTS categories (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(255) NOT NULL,
    hex_color VARCHAR(7) NOT NULL,
    nerd_icon VARCHAR(10) NOT NULL
);
