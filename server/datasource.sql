-- auto-generated definition
create table id_token
(
    id       int auto_increment
        primary key,
    biz_type varchar(36)             not null,
    remake   varchar(100) default '' not null,
    token    varchar(128)            not null,
    constraint unique_biz_key
        unique (biz_type)
);



-- auto-generated definition
create table id_info
(
    id       int auto_increment
        primary key,
    biz_type varchar(36)   not null,
    step     int           not null comment '每次获取号段的长度,不代表个数',
    incr     int default 1 not null comment 'id每次增加的增量',
    max_id   int default 1 not null,
    version  int default 0 not null,
    constraint unique_biz_key
        unique (biz_type)
);

