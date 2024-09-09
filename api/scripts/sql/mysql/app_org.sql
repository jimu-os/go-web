-- 组织
drop table if exists app_org;
create table app_org
(
    id          bigint primary key,
    pid         varchar(30) comment '上级组织',
    name        varchar(30) comment '组织名',
    create_time timestamp(0) not null default now()
) comment '组织表';


-- 初始化组织
insert into app_org(id, pid, name)
values (1, '', '根组织');

-- 用户
drop table if exists user;
create table if not exists user
(
    id          bigint primary key comment 'id',
    name        varchar(50)  not null comment '用户昵称',
    account     varchar(50)  not null unique comment '用户注册账号',
    phone       varchar(11)  not null default '' comment '手机',
    email       varchar(100) not null default '' comment '邮箱',
    password    varchar(100) not null default '' comment '用户密码',
    picture     varchar(500) not null default '' comment '用户头像',
    gender      int          not null default 0 comment '性别',
    age         int          not null default 0 comment '年龄',
    status      boolean      not null default true comment '状态',
    create_time timestamp(0)          default now() not null
);

create index user_account_key on user (account);
create index user_phone_key on user (phone);
create index user_email_key on user (email);
create index ape on user (account, phone, email);


-- 初始化root用户
insert into user(id, name, picture, account, password, gender, age, phone, email)
values (1, 'root', 'https://im-1252940994.cos.ap-nanjing.myqcloud.com/go.jpg', 'root',
        'PXeFkSq39sRSVXsSaxLWUUWDt45I8tw9mgcY8GE3B/r3VylOko0q727qPChy+uibqcFuCy6w67ruaQ3AyHHeDg==', 0, 12,
        '15527574117', '1219449282@qq.com');

-- 角色
drop table if exists role;
create table role
(
    id          bigint primary key,
    name        varchar(30),
    role_key    varchar(50)  not null,
    status      boolean      not null default true,
    create_time timestamp(0) not null default now()
);


-- 初始化角色
insert into role(id, name, role_key)
values (1, '超级管理员', 'root');
insert into role(id, name, role_key)
values (2, '管理员', 'admin');
insert into role(id, name, role_key)
values (3, '普通用户', 'normal');


