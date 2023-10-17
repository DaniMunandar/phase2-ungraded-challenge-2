-- Membuat tabel Heroes
CREATE TABLE Heroes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    universe VARCHAR(255) NOT NULL,
    skill VARCHAR(255) NOT NULL,
    imageURL VARCHAR(255) NOT NULL
);

-- Membuat tabel Villain
CREATE TABLE Villain (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    universe VARCHAR(255) NOT NULL,
    imageURL VARCHAR(255) NOT NULL
);

-- Memasukkan contoh data ke dalam tabel Heroes
INSERT INTO Heroes (name, universe, skill, imageURL) VALUES
    ('Iron Man', 'Marvel', 'Genius inventor and suit of armor', 'ironman.jpg'),
    ('Captain America', 'Marvel', 'Super Soldier', 'captainamerica.jpg'),
    ('Thor', 'Marvel', 'God of Thunder', 'thor.jpg');

-- Memasukkan contoh data ke dalam tabel Villain
INSERT INTO Villain (name, universe, imageURL) VALUES
    ('Loki', 'Marvel', 'loki.jpg'),
    ('Red Skull', 'Marvel', 'redskull.jpg'),
    ('Thanos', 'Marvel', 'thanos.jpg');
