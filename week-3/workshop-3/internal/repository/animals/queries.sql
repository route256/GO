-- name: ListAnimals :many
SELECT * FROM animals;

-- name: GetAnimal :one
SELECT * FROM animals WHERE id = $1;

-- name: Insert :exec
insert into animals (nickname, birthday, weight) VALUES ($1, $2, $3); 