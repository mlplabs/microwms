create table if not exists public.users
(
    id   bigserial
        constraint users_pk
            primary key,
    name varchar(100) default ''::character varying not null
);

create table if not exists public.manufacturers
(
    id bigserial
        constraint manufacturers_pk
            primary key,
    name varchar(250) not null
);

create table if not exists public.barcodes
(
    id           bigserial
        constraint barcodes_pk
            primary key,
    name         text        default ''::text              not null,
    barcode_type integer     default 0                     not null,
    owner_id     integer     default 0                     not null,
    owner_ref    varchar(36) default ''::character varying not null
);

comment on column public.barcodes.owner_id is 'id подчиненного элемента';
comment on column public.barcodes.owner_ref is 'таблица владельца шк';

create table if not exists public.products
(
    id              bigserial
        constraint products_pk
            primary key,
    name            text                                      not null,
    manufacturer_id integer     default 0                     not null,
    item_number     varchar(50) default ''::character varying not null
);

comment on column public.products.item_number is 'артикул';

create table if not exists public.zones
(
    id        bigserial
        constraint zones_pk
            primary key,
    name      varchar(50) default ''::character varying,
    parent_id integer,
    zone_type smallint
);

create table public.cells
(
    id              bigserial
        constraint cells_pk
            primary key,
    name            varchar(50)   default ''::character varying not null,
    whs_id          integer       default 0                     not null,
    zone_id         integer       default 0                     not null,
    section_id      integer       default 0                     not null,
    passage_id      integer       default 0                     not null,
    rack_id         integer       default 0                     not null,
    floor           integer       default 0                     not null,
    number          integer       default 0                     not null,
    sz_length       integer       default 0                     not null,
    sz_width        integer       default 0                     not null,
    sz_height       integer       default 0                     not null,
    sz_volume       numeric(8, 3) default 0                     not null,
    sz_uf_volume    numeric(8, 3) default 0                     not null,
    sz_weight       numeric(8, 3) default 0                     not null,
    is_size_free    boolean       default false                 not null,
    is_weight_free  boolean       default false                 not null,
    not_allowed_in  boolean       default false                 not null,
    not_allowed_out boolean       default false                 not null,
    is_service      boolean       default false                 not null
);

create table if not exists public.warehouses
(
    id   bigserial
        constraint warehouses_pk
            primary key,
    name varchar(100) default ''::character varying not null,
    address text default '' not null
);
