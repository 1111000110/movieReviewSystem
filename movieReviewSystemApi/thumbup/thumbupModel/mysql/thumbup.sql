create table thumbUp(
                         thumb_up_id bigint AUTO_INCREMENT,
                         user_id bigint NOT NULL DEFAULT 0,
                         base_id bigint NOT NULL DEFAULT 0,
                         val bigint NOT NULL DEFAULT 0,
                         create_at bigint  NOT NULL DEFAULT '0',
                         update_at bigint NOT NULL DEFAULT '0',
                         PRIMARY KEY (thumb_up_id)
);