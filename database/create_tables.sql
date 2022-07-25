CREATE TABLE quotes (
    quoteID INT,
    authorID INT, 
    quote VARCHAR(200), 
    last_used DATETIME
);

CREATE TABLE authors (
    authorID INT,
    name VARCHAR(50),
    font VARCHAR(50)
);

CREATE TABLE backgrounds (
    pictureID INT,
    name VARCHAR(50),
    image BLOB
);