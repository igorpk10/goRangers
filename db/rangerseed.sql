-- 🚨 Certifique-se de rodar este script após criar as tabelas

-- 🧑‍🎤 Rangers
INSERT INTO ranger (
    name, morphed_name, color, role, team, season,
    zord, weapon, morpher, actor_name, planet_origin,
    first_appearance, age, is_leader, status, bio
)
VALUES
-- Jason Lee Scott - Red Ranger
('Jason Lee Scott', 'Red Ranger', 'Red', 'Leader', 'Mighty Morphin Power Rangers', 'Season 1',
 'Tyrannosaurus Dinozord', 'Power Sword', 'Dino Morpher', 'Austin St. John', 'Earth',
 '1993-08-28', 18, TRUE, 'Ativo', 'Jason é o líder dos Power Rangers originais e mestre em artes marciais.'),

-- Zack Taylor - Black Ranger
('Zack Taylor', 'Black Ranger', 'Black', 'Dance Fighter', 'Mighty Morphin Power Rangers', 'Season 1',
 'Mastodon Dinozord', 'Power Axe', 'Dino Morpher', 'Walter Emanuel Jones', 'Earth',
 '1993-08-28', 18, FALSE, 'Ativo', 'Zack é o divertido e energético Ranger preto, especialista em danças de combate.'),

-- Kimberly Ann Hart - Pink Ranger
('Kimberly Ann Hart', 'Pink Ranger', 'Pink', 'Gymnast', 'Mighty Morphin Power Rangers', 'Season 1',
 'Pterodactyl Dinozord', 'Power Bow', 'Dino Morpher', 'Amy Jo Johnson', 'Earth',
 '1993-08-28', 18, FALSE, 'Ativo', 'Kimberly é uma ginasta talentosa e a primeira Ranger rosa.'),

-- Trini Kwan - Yellow Ranger
('Trini Kwan', 'Yellow Ranger', 'Yellow', 'Martial Artist', 'Mighty Morphin Power Rangers', 'Season 1',
 'Saber-Toothed Tiger Dinozord', 'Power Daggers', 'Dino Morpher', 'Thuy Trang', 'Earth',
 '1993-08-28', 18, FALSE, 'Ativo', 'Trini é inteligente, calma e possui grande domínio das artes marciais.'),

-- Billy Cranston - Blue Ranger
('Billy Cranston', 'Blue Ranger', 'Blue', 'Tech Expert', 'Mighty Morphin Power Rangers', 'Season 1',
 'Triceratops Dinozord', 'Power Lance', 'Dino Morpher', 'David Yost', 'Earth',
 '1993-08-28', 18, FALSE, 'Ativo', 'Billy é o gênio do grupo, responsável por muitas das invenções técnicas.'),

-- Tommy Oliver - Green Ranger
('Tommy Oliver', 'Green Ranger', 'Green', 'Lone Warrior', 'Mighty Morphin Power Rangers', 'Season 1',
 'Dragonzord', 'Dragon Dagger', 'Dragon Morpher', 'Jason David Frank', 'Earth',
 '1993-10-05', 18, FALSE, 'Ativo', 'Tommy começou como inimigo dos Rangers, mas depois se tornou um dos membros mais poderosos.')
;


-- 🖼️ Pictures

-- Jason
INSERT INTO picture (ranger_id, url, description)
VALUES
(1, 'https://static.wikia.nocookie.net/powerrangers/images/7/7a/MMPR_Red_Ranger.png', 'Jason transformado como Red Ranger');

-- Zack
INSERT INTO picture (ranger_id, url, description)
VALUES
(2, 'https://static.wikia.nocookie.net/powerrangers/images/4/49/MMPR_Black_Ranger.png', 'Zack em forma de Black Ranger');

-- Kimberly
INSERT INTO picture (ranger_id, url, description)
VALUES
(3, 'https://static.wikia.nocookie.net/powerrangers/images/2/2d/MMPR_Pink_Ranger.png', 'Kimberly como Pink Ranger');

-- Trini
INSERT INTO picture (ranger_id, url, description)
VALUES
(4, 'https://static.wikia.nocookie.net/powerrangers/images/3/37/MMPR_Yellow_Ranger.png', 'Trini em uniforme de Yellow Ranger');

-- Billy
INSERT INTO picture (ranger_id, url, description)
VALUES
(5, 'https://static.wikia.nocookie.net/powerrangers/images/6/64/MMPR_Blue_Ranger.png', 'Billy com seu traje azul');

-- Tommy
INSERT INTO picture (ranger_id, url, description)
VALUES
(6, 'https://static.wikia.nocookie.net/powerrangers/images/5/59/MMPR_Green_Ranger.png', 'Tommy como Green Ranger com a Dragon Dagger');
