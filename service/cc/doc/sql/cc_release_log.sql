create table cc_release_log
(
    id           int auto_increment,
    ip           varchar(32) default '' COMMENT '客户端ip',
    kv_id        bigint      default 0 COMMENT 'cc_kv id',
    app_id       bigint      default 0 COMMENT 'app_id',
    cluster_id   bigint      default 0 COMMENT 'cluster_id',
    kv_revision_id   bigint      default 0 COMMENT 'cc_kv_revision id',
    release_time timestamp COMMENT '下发配置时间',
    create_time  timestamp   default CURRENT_TIMESTAMP COMMENT '',
    update_time  timestamp null  default null on UPDATE CURRENT_TIMESTAMP COMMENT '',
    PRIMARY KEY (`id`),
    index idx_create_time (`create_time`),
    index idx_kv_id(`kv_id`)
) comment '下发日志';