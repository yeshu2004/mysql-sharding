CREATE TABLE album (
    id VARCHAR(36) NOT NULL,
    title VARCHAR(128) NOT NULL,
    artist VARCHAR(128) NOT NULL,
    price DECIMAL(5,2) NOT NULL,
    PRIMARY KEY (id)
) 

-- INSERT INTO album
--     (title, artist, price)
--     VALUES
--     ('Kind of Blue', 'Miles Davis', 45.99),
--     ('Time Out', 'Dave Brubeck', 39.99),
--     ('A Love Supreme', 'John Coltrane', 59.99),
--     ('Mingus Ah Um', 'Charles Mingus', 29.99);