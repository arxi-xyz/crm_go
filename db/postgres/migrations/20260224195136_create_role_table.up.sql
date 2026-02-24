CREATE TABLE IF NOT EXISTS roles
(
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    uuid       UUID        NOT NULL DEFAULT gen_random_uuid(),

    title      TEXT        NOT NULL,

    parent_id  bigint,
    constraint fk_role foreign key (parent_id) references roles (id),

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX roles_uuid_unique
    ON roles (uuid);

create table if not exists model_has_role
(
    model_type text,
    model_id   integer,
    role_id    bigint,

    constraint fk_role foreign key (role_id) references roles (id)
);


CREATE INDEX has_role_index
    ON model_has_role (model_type, model_id);
