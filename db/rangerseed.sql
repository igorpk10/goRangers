-- GENRES
INSERT INTO genres (name, image_url, color_theme)
VALUES 
('Super Sentai', 'https://example.com/super-sentai.png', 'red');

-- FRANCHISE
INSERT INTO franchises (name, genre_id, youtube_channel, description, logo_url)
VALUES (
  'Power Rangers',
  1,
  'https://www.youtube.com/@powerrangers',
  'Power Rangers is a long-running American entertainment franchise that adapts Super Sentai shows from Japan.',
  'https://example.com/power-rangers-logo.png'
);

-- SEASON
INSERT INTO seasons (franchise_id, title, subtitle, description, image_url, release_year, num_episodes, theme_song_url, opening_credits_url)
VALUES (
  1,
  'Mighty Morphin Power Rangers',
  'Season 1',
  'The original season where five teenagers with attitude are chosen to fight evil.',
  'https://example.com/mmpr-season1.jpg',
  1993,
  60,
  'https://example.com/mmpr-theme.mp3',
  'https://example.com/opening.mp4'
);

-- RANGERS
INSERT INTO rangers (name, actor_name, color, bio, trivia, status)
VALUES 
('Jason Lee Scott', 'Austin St. John', 'Red', 'Leader of the original team.', 'He is the original Red Ranger.', 'Alive'),
('Kimberly Hart', 'Amy Jo Johnson', 'Pink', 'Gymnast and musician.', 'She played the Pink Ranger and was a fan favorite.', 'Alive'),
('Zack Taylor', 'Walter Emanuel Jones', 'Black', 'Fun-loving and energetic.', 'Known for his Hip Hop Kido fighting style.', 'Alive'),
('Trini Kwan', 'Thuy Trang', 'Yellow', 'Calm and intelligent.', 'Trini was the original Yellow Ranger.', 'Deceased'),
('Billy Cranston', 'David Yost', 'Blue', 'Tech genius of the group.', 'He developed many gadgets for the team.', 'Alive'),
('Tommy Oliver', 'Jason David Frank', 'Green', 'The original Green Ranger.', 'Started as a villain, became the most iconic Ranger.', 'Deceased');

-- SEASON ↔ RANGERS
INSERT INTO season_rangers (season_id, ranger_id, role, morphing_call)
VALUES 
(1, 1, 'Leader', 'It’s Morphin Time! Tyrannosaurus!'),
(1, 2, 'Main', 'It’s Morphin Time! Pterodactyl!'),
(1, 3, 'Main', 'It’s Morphin Time! Mastodon!'),
(1, 4, 'Main', 'It’s Morphin Time! Sabertooth Tiger!'),
(1, 5, 'Main', 'It’s Morphin Time! Triceratops!'),
(1, 6, 'Guest', 'It’s Morphin Time! Dragonzord!');

-- ZORDS
INSERT INTO zords (season_id, name, type, description, image_url)
VALUES
(1, 'Tyrannosaurus Dinozord', 'Dinosaur', 'Jason’s Zord', 'https://example.com/tyrannosaurus.jpg'),
(1, 'Pterodactyl Dinozord', 'Dinosaur', 'Kimberly’s Zord', 'https://example.com/pterodactyl.jpg'),
(1, 'Mastodon Dinozord', 'Dinosaur', 'Zack’s Zord', 'https://example.com/mastodon.jpg'),
(1, 'Sabertooth Tiger Dinozord', 'Dinosaur', 'Trini’s Zord', 'https://example.com/sabertooth.jpg'),
(1, 'Triceratops Dinozord', 'Dinosaur', 'Billy’s Zord', 'https://example.com/triceratops.jpg'),
(1, 'Dragonzord', 'Mythical', 'Tommy’s Zord', 'https://example.com/dragonzord.jpg');

-- VILLAINS
INSERT INTO villains (season_id, name, description, main_villain, actor_name, image_url)
VALUES 
(1, 'Rita Repulsa', 'Evil space witch trying to conquer Earth.', TRUE, 'Machiko Soga (Japanese) / Barbara Goodson (voice)', 'https://example.com/rita.jpg'),
(1, 'Goldar', 'Rita’s loyal warrior.', FALSE, 'Kerrigan Mahan (voice)', 'https://example.com/goldar.jpg');

-- EPISODES (only a few examples)
INSERT INTO episodes (season_id, title, synopsis, episode_number, air_date, is_key_episode, thumbnail_url)
VALUES 
(1, 'Day of the Dumpster', 'The Rangers are chosen to fight Rita Repulsa.', 1, '1993-08-28', TRUE, 'https://example.com/ep1.jpg'),
(1, 'Green with Evil: Part 1', 'Tommy becomes the evil Green Ranger.', 17, '1993-10-05', TRUE, 'https://example.com/ep17.jpg'),
(1, 'Green with Evil: Part 5', 'The Rangers free Tommy from evil.', 21, '1993-10-09', TRUE, 'https://example.com/ep21.jpg');

-- RANGER PHOTOS
INSERT INTO ranger_photos (ranger_id, photo_url, caption)
VALUES 
(1, 'https://example.com/jason1.jpg', 'Jason in Red Ranger suit'),
(2, 'https://example.com/kimberly1.jpg', 'Kimberly in action');

-- GENERATED IMAGES
INSERT INTO ranger_generated_images (ranger_id, image_url, description)
VALUES 
(1, 'https://ai-generated.com/jason-ai1.png', 'AI styled Jason in Red Ranger armor'),
(6, 'https://ai-generated.com/tommy-dragonzord.png', 'Green Ranger summoning Dragonzord');

-- WEAPONS
INSERT INTO weapons (name, ranger_id, description, image_url)
VALUES 
('Power Sword', 1, 'Red Ranger’s primary weapon', 'https://example.com/power-sword.jpg'),
('Power Bow', 2, 'Pink Ranger’s ranged weapon', 'https://example.com/power-bow.jpg'),
('Dragon Dagger', 6, 'Used to summon the Dragonzord', 'https://example.com/dragon-dagger.jpg');

-- SEASON TRIVIA
INSERT INTO season_trivia (season_id, fact)
VALUES 
(1, 'The show used footage from the Japanese series Kyōryū Sentai Zyuranger.'),
(1, 'The Green Ranger was originally meant to appear in only 5 episodes.');
