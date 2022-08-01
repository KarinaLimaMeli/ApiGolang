-- Mostar todos os registros da tavbela de filmes
select * from movies;

-- Mostrar o nome, sobrenome e calssificação de todos os atores
select first_name, last_name, rating 
from actors;
-- Mostrar o titulo de todas as series e use alias para que tanto o nome da tabela quanto o campo esteja em portugues.
select séries.title Título
 from series séries;

-- Mostrar o nome e o sobrenome dos atores cuja classificação é superior a 7,5
select first_name, last_name, rating
from actors
where rating > 7.5;

-- Mostar o título dos filmes, a classificação e os prêmios dos filmes com classificação
-- Superior a 7,5 e com mais de dois prêmios
select title, rating, awards
from movies
where rating > 7.5 AND awards > 2;

-- Mostraro titulo e a classificacao ordenados por classificação em ordem crescente 
select title, rating
from movies
order by rating desc;

-- Mostrar Titulo dos 3 primeiros filmes
select title
from movies
limit 3;
-- Mostrar os 5 melhores filmes com classificação mais alta
select title, rating
from movies
order by rating desc
limit 5;

-- Listar os 10 primeiros atores
select first_name, last_name
from actors
limit 10;
-- Mostar o título e a classificação de todos os filmes cujo título é Toy Story
select title
from movies
where title = 'Toy Story';
-- Mostar todos os atores cujos os nomes começam com Sam.
select title, rating
select first_name, last_name
from actors
where first_name like 'Sam%';

