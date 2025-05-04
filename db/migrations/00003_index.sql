-- +goose Up
-- +goose StatementBegin
create index url_index on urls using hash (short_url);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index url_index;
-- +goose StatementEnd
