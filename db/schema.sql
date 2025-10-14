-- GENRES
CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    image_url TEXT,
    color_theme VARCHAR(50)
);

-- FRANCHISES
CREATE TABLE franchises (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    genre_id INT REFERENCES genres(id),
    youtube_channel TEXT,
    description TEXT,
    logo_url TEXT
);

-- SEASONS
CREATE TABLE seasons (
    id SERIAL PRIMARY KEY,
    franchise_id INT REFERENCES franchises(id),
    title VARCHAR(150) NOT NULL,
    subtitle VARCHAR(255),
    description TEXT,
    image_url TEXT,
    release_year INT,
    num_episodes INT,
    theme_song_url TEXT,
    opening_credits_url TEXT
);

-- EPISODES
CREATE TABLE episodes (
    id SERIAL PRIMARY KEY,
    season_id INT REFERENCES seasons(id),
    title VARCHAR(150) NOT NULL,
    synopsis TEXT,
    episode_number INT,
    air_date DATE,
    is_key_episode BOOLEAN DEFAULT FALSE,
    thumbnail_url TEXT
);

-- RANGERS
CREATE TABLE rangers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    actor_name VARCHAR(150),
    color VARCHAR(50),
    bio TEXT,
    trivia TEXT,
    debut_episode_id INT REFERENCES episodes(id),
    status VARCHAR(50)
);

-- RANGER PHOTOS
CREATE TABLE ranger_photos (
    id SERIAL PRIMARY KEY,
    ranger_id INT REFERENCES rangers(id),
    photo_url TEXT NOT NULL,
    caption TEXT
);

-- GENERATED IMAGES (MCP)
CREATE TABLE ranger_generated_images (
    id SERIAL PRIMARY KEY,
    ranger_id INT REFERENCES rangers(id),
    image_url TEXT NOT NULL,
    description TEXT
);

-- SEASON ↔ RANGERS (Many-to-Many)
CREATE TABLE season_rangers (
    season_id INT REFERENCES seasons(id),
    ranger_id INT REFERENCES rangers(id),
    PRIMARY KEY (season_id, ranger_id),
    role VARCHAR(100),
    morphing_call TEXT
);

-- ZORDS
CREATE TABLE zords (
    id SERIAL PRIMARY KEY,
    season_id INT REFERENCES seasons(id),
    name VARCHAR(150) NOT NULL,
    type VARCHAR(100),
    description TEXT,
    image_url TEXT
);

-- VILLAINS
CREATE TABLE villains (
    id SERIAL PRIMARY KEY,
    season_id INT REFERENCES seasons(id),
    name VARCHAR(150) NOT NULL,
    description TEXT,
    main_villain BOOLEAN DEFAULT FALSE,
    actor_name VARCHAR(150),
    image_url TEXT
);


-- WEAPONS
CREATE TABLE weapons (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150),
    ranger_id INT REFERENCES rangers(id),
    description TEXT,
    image_url TEXT
);

-- TRIVIA PER SEASON
CREATE TABLE season_trivia (
    id SERIAL PRIMARY KEY,
    season_id INT REFERENCES seasons(id),
    fact TEXT
);

-- CAMEOS
CREATE TABLE cameos (
    id SERIAL PRIMARY KEY,
    ranger_id INT REFERENCES rangers(id),
    episode_id INT REFERENCES episodes(id),
    cameo_type VARCHAR(100)
);
