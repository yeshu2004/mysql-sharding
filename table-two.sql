CREATE TABLE album (
    id INT AUTO_INCREMENT,
    title VARCHAR(128) NOT NULL,
    artist VARCHAR(128) NOT NULL,
    price DECIMAL(5,2) NOT NULL,
    PRIMARY KEY (id)
) AUTO_INCREMENT=5;

INSERT INTO album
    (title, artist, price)
    VALUES
    ('Kind of Blue', 'Miles Davis', 45.99),
    ('Time Out', 'Dave Brubeck', 39.99),
    ('A Love Supreme', 'John Coltrane', 59.99),
    ('Mingus Ah Um', 'Charles Mingus', 29.99);