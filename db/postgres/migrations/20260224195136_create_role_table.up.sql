CREATE TABLE IF NOT EXISTS roles
(
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    uuid       UUID        NOT NULL DEFAULT gen_random_uuid(),

    title      TEXT        NOT NULL,

    parent_id  bigint,
    constraint fk_roles_parent foreign key (parent_id) references roles (id),

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

    constraint fk_model_role_role foreign key (role_id) references roles (id)
);


CREATE INDEX has_role_index
    ON model_has_role (model_type, model_id);

insert into roles (title, parent_id)
values ('super_admin', null);
insert into roles (title, parent_id)
values ('admin', 1);

insert into model_has_permission (model_type, model_id, permission_id)
select 'roles', r.id, p.id
from roles r
         cross join permissions p
where r.title = 'super_admin'
on conflict do nothing;
