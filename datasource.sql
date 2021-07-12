# 业务方信息
create table id_token
(
    id       int auto_increment
        primary key,
    biz_type varchar(36)             not null,
    remake   varchar(100) default '' not null comment '备注',
    token    varchar(128)            not null,
    constraint unique_biz_key
        unique (biz_type)
);


# 业务方id号段信息
create table id_info
(
    id        int auto_increment
        primary key,
    biz_type  varchar(36)   not null,
    step      int           not null comment '每次获取号段的长度,不代表个数',
    incr      int default 1 not null comment 'id每次增加的增量',
    remainder int           not null,
    max_id    int default 1 not null,
    version   int default 0 not null,
    constraint unique_biz_key
        unique (biz_type)
);

