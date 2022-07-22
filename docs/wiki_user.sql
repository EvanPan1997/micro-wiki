create table wiki_user
(
    data_id              VARCHAR(64) primary key,

    user_id              VARCHAR(64),
    username             VARCHAR(128),
    password             VARCHAR(512),
    role_id              VARCHAR(64),
    status               CHAR(1),
    is_locked            CHAR(1),
    last_login_time      VARCHAR(14),
    last_logout_time     VARCHAR(14),
    login_ip             VARCHAR(128),
    pwd_error_count      int,
    total_pwd_error      int,
    last_change_pwd_time VARCHAR(14),
    access_token         VARCHAR(512),

    create_user          VARCHAR(64),
    create_at            VARCHAR(14),
    change_user          VARCHAR(64),
    change_at            VARCHAR(14),
    delete_user          VARCHAR(64),
    delete_at            VARCHAR(14),

    remarks              VARCHAR(2000),
    reserve1             VARCHAR(128),
    reserve2             VARCHAR(128),
    reserve3             VARCHAR(128),
    reserve4             VARCHAR(128),
    reserve5             VARCHAR(128)
);