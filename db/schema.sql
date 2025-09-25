CREATE TABLE ranger (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    morphed_name VARCHAR(100),
    color VARCHAR(50),
    role VARCHAR(100),
    team VARCHAR(100),
    season INT,
    zord VARCHAR(100),
    weapon VARCHAR(100),
    morpher VARCHAR(100),
    actor_name VARCHAR(100),
    planet_origin VARCHAR(100),
    first_appearance DATE,
    age INT,
    is_leader BOOLEAN DEFAULT FALSE,
    status VARCHAR(50) DEFAULT 'Ativo',
    bio TEXT
);


CREATE TABLE picture (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    ranger_id INT REFERENCES ranger(id) ON DELETE CASCADE,
    url VARCHAR(255) NOT NULL,
    description TEXT
);