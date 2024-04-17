CREATE TABLE `user`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '사용자 식별자',
    `name`     varchar(20) NOT NULL COMMENT '사용자 이름',
    `password` VARCHAR(80) NOT NULL COMMENT '패스워드 해시',
    `role`     VARCHAR(80) NOT NULL COMMENT '역할',
    `created`  DATETIME(6) NOT NULL COMMENT '작성시각',
    `modified` DATETIME(6) NOT NULL COMMENT '수정시각',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='사용자';

CREATE TABLE `task`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '할일 식별자',
    `title`    VARCHAR(128) NOT NULL COMMENT '할일 제목',
    `status`   VARCHAR(20)  NOT NULL COMMENT '할일 상태',
    `created`  DATETIME(6) NOT NULL COMMENT '작성시각',
    `modified` DATETIME(6) NOT NULL COMMENT '수정시각',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='할일';
