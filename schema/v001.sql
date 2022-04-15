create table docker
(
    id         varchar(64)  not null
        primary key,
    host       varchar(200) null,
    user_name  varchar(64)  null,
    secret     text         null,
    name       varchar(64)  null,
    created_at bigint       null,
    updated_at bigint       null,
    deleted_at bigint       null,
    created_by varchar(64)  null,
    updated_by varchar(64)  null,
    deleted_by varchar(64)  null,
    tenant_id  varchar(64)  null
);

create table function
(
    id         varchar(64)  not null
        primary key,
    group_name varchar(200) null,
    project    varchar(200) null,
    version    varchar(200) null,
    language   varchar(200) null,
    status     varchar(200) null,
    env        text         null,
    created_at bigint       null,
    updated_at bigint       null,
    deleted_at bigint       null,
    created_by varchar(64)  null,
    updated_by varchar(64)  null,
    deleted_by varchar(64)  null,
    tenant_id  varchar(64)  null
);

create table git
(
    id         varchar(64)  not null
        primary key,
    host       varchar(200) null,
    token      text         null,
    name       varchar(200) null,
    created_at bigint       null,
    updated_at bigint       null,
    deleted_at bigint       null,
    created_by varchar(64)  null,
    updated_by varchar(64)  null,
    deleted_by varchar(64)  null,
    tenant_id  varchar(64)  null
);

