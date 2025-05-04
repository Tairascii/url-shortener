-- +goose Up
-- +goose StatementBegin
create table urls (
    id bigint primary key,
    short_url text,
    long_url text
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table urls;
-- +goose StatementEnd
