
-- name: GetFullUrl :one
select long_url from urls where short_url = $1;

-- name: GenerateID :one
insert into urls (id) values (nextval('url_id_manual_seq') + $1 + $2) returning id;

-- name: SetUrl :exec
update urls set long_url = $1, short_url = $2 where id = $3;

