CREATE TABLE album (
    id INT AUTO_INCREMENT ,
    title VARCHAR(128) NOT NULL,
    artist VARCHAR(128) NOT NULL,
    price DECIMAL(5,2) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO album
    (title, artist, price)
    VALUES
    ('Blue Train', 'John Coltrane', 56.99),
    ('Giant Steps', 'John Coltrane', 63.99),
    ('Jeru', 'Gerry Mulligan', 17.99),
    ('Sarah Vaughan', 'Sarah Vaughan', 34.98);