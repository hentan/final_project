CREATE TABLE authors(
                        ID serial PRIMARY KEY,
                        Name_author varchar(100),
                        Surname_author varchar(100),
                        Biography text,
                        Birthday date
);

CREATE TABLE books(
                      ID serial PRIMARY KEY,
                      Name_book varchar(300),
                      Author_ID int,
                      Year_book int,
                      ISBN varchar(50)
);

ALTER TABLE books
    ADD CONSTRAINT fk_books_authors FOREIGN KEY(Author_ID)
        REFERENCES authors(ID);

INSERT INTO authors(Name_author, Surname_author, Biography, Birthday)
VALUES
    (
        'Leo',
        'Tolstoy',
        'Count Lev Nikolayevich Tolstoy was a Russian writer...',
        '1828-09-09'
    ),
    (
        'Fyodor',
        'Dostoevsky',
        'Fyodor Mikhailovich Dostoevsky was a Russian novelist...',
        '1821-11-11'
    );

INSERT INTO books(Name_book, Author_ID, Year_book, ISBN)
VALUES
    ('War and Peace', 1, 1869, '978-5-389-06256-6'),
    ('The Idiot', 2, 1868, '978-1533695840');
