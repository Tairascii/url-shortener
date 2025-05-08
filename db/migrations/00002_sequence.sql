-- +goose Up
-- +goose StatementBegin
create sequence url_id_manual_seq increment 10000 start 10000;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop sequence url_id_manual_seq;
-- +goose StatementEnd
