-- +goose Up
-- +goose StatementBegin
create sequence url_id_manual_seq increment 1000 start 1000;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop sequence url_id_manual_seq;
-- +goose StatementEnd
