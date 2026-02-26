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

CREATE UNIQUE INDEX IF NOT EXISTS model_has_permission_unique
    ON model_has_permission (model_type, model_id, permission_id);

insert into permissions (title, unique_key)
values ('all', 'all');

insert into permissions (title, unique_key)
values ('create_profile', 'create_profile');
insert into permissions (title, unique_key)
values ('update_profile', 'update_profile');
insert into permissions (title, unique_key)
values ('delete_profile', 'delete_profile');
insert into permissions (title, unique_key)
values ('view_profile', 'view_profile');