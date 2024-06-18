create table authors(
	ID serial primary key,
	Name_author varchar(100),
	Sirname_author varchar(100),
	Biography varchar,
	Birthday date
);

create table books(
	ID serial primary key,
	Name_book varchar(300),
	Author_ID int,
	Year_book int,
	ISBN varchar(50)
);

alter table books
add constraint fk_books_authors foreign key(Author_ID)
references authors(ID);

insert into authors(Name_author, Sirname_author, Biography, Birthday)
values(
    'Leo',
    'Tolstoy',
    'Count Lev Nikolayevich Tolstoy[note 1] (/ˈtoʊlstɔɪ, ˈtɒl-/;[1] Russian: Лев Николаевич Толстой,[note 2] IPA: [ˈlʲef nʲɪkɐˈla(j)ɪvʲɪtɕ tɐlˈstoj] ⓘ; 9 September [O.S. 28 August] 1828 – 20 November [O.S. 7 November] 1910),[2] usually referred to in English as Leo Tolstoy, was a Russian writer. He is regarded as one of the greatest and most influential authors of all time.[3][4] He received nominations for the Nobel Prize in Literature every year from 1902 to 1906 and for the Nobel Peace Prize in 1901, 1902, and 1909',
    '1828-09-09'
),
(
    'Fyodor',
    'Dostoevsky',
    'Fyodor Mikhailovich Dostoevsky[a] (UK: /ˌdɒstɔɪˈɛfski/,[1] US: /ˌdɒstəˈjɛfski, ˌdʌs-/;[2] Russian: Фёдор Михайлович Достоевский[b], romanized: Fyodor Mikhaylovich Dostoyevskiy, IPA: [ˈfʲɵdər mʲɪˈxajləvʲɪdʑ dəstɐˈjefskʲɪj] ⓘ; 11 November 1821 – 9 February 1881[3][c]), sometimes transliterated as Dostoyevsky, was a Russian novelist, short story writer, essayist and journalist. Numerous literary critics regard him as one of the greatest novelists in all of world literature, as many of his works are considered highly influential masterpieces',
    '1821-11-11'
);

insert into books(Name_book, Author_ID, Year_book, ISBN)
values
('War and Peace', 1, 1869, '978-5-389-06256-6'),
('the Idiot', 2, 1868, '978-1533695840');

