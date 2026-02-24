CREATE TABLE IF NOT EXISTS permissions
(
    id         integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    uuid       UUID NOT NULL DEFAULT gen_random_uuid(),

    title      TEXT NOT NULL,
    unique_key TEXT NOT NULL unique
);

create table if not exists model_has_permission
(
    model_type    text,
    model_id      integer,
    permission_id integer,

    constraint fk_permission foreign key (permission_id) references permissions (id)
);

CREATE INDEX has_permission_index
    ON model_has_permission (model_type, model_id);