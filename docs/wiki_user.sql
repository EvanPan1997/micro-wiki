create table wiki_user
(
    data_id              VARCHAR(64) primary key,

    user_id              VARCHAR(64),
    username             VARCHAR(128),
    password             VARCHAR(512),
    role_id              VARCHAR(64),
    status               CHAR(1),
    is_locked            CHAR(1),
    last_login_time      DATETIME,
    last_logout_time     DATETIME,
    login_ip             VARCHAR(128),
    pwd_error_count      int,
    total_pwd_error      int,
    last_change_pwd_time DATETIME,
    access_token         VARCHAR(512),

    create_user          VARCHAR(64),
    create_at            DATETIME,
    changeUser           VARCHAR(64),
    change_at            DATETIME,
    delete_user          VARCHAR(64),
    delete_at            DATETIME,

    remarks              VARCHAR(2000),
    reserve1             VARCHAR(128),
    reserve2             VARCHAR(128),
    reserve3             VARCHAR(128),
    reserve4             VARCHAR(128),
    reserve5             VARCHAR(128)
);