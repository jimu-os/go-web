drop table if exists app_dictionary;
create table app_dictionary
(
    id          bigint primary key auto_increment comment '字典id',
    type        varchar(30) not null comment '字典类别',
    status      int         not null comment '字典状态值',
    name        varchar(30) not null comment '字典名称',
    create_time timestamp(0) default now() comment '创建时间'
);


-- 初始化字典表
insert into app_dictionary(type, status, name)
values ('notify', 1, 'info');

insert into app_dictionary(type, status, name)
values ('notify', 2, 'success');

insert into app_dictionary(type, status, name)
values ('notify', 3, 'warning');

insert into app_dictionary(type, status, name)
values ('notify', 4, 'error');

insert into app_dictionary(type, status, name)
values ('notify', 5, '已处理');

insert into app_dictionary(type, status, name)
values ('notify', 6, '未处理');

insert into app_dictionary(type, status, name)
values ('user', 0, '男');

insert into app_dictionary(type, status, name)
values ('user', 1, '女');

insert into app_dictionary(type, status, name)
values ('tool', 1, '路由工具');

insert into app_dictionary(type, status, name)
values ('tool', 2, '窗口工具');

insert into app_dictionary(type, status, name)
values ('tool_position', 1, '左侧');

insert into app_dictionary(type, status, name)
values ('tool_position', 2, '右侧');

insert into app_dictionary(type, status, name)
values ('chat_knowledge_file', 1, '文件夹');

insert into app_dictionary(type, status, name)
values ('chat_knowledge_file', 2, 'docx');

insert into app_dictionary(type, status, name)
values ('chat_knowledge_file', 3, 'pdf');

insert into app_dictionary(type, status, name)
values ('chat_knowledge_file', 4, 'excel');

insert into app_dictionary(type, status, name)
values ('chat_knowledge_file', 5, 'txt');


-- 消息表
drop table if exists app_notify;
create table app_notify
(
    id          bigint primary key auto_increment,
    pub_id      varchar(30)  not null comment '生产者',
    sub_id      varchar(30)  not null comment '消费者',
    title       varchar(50)  not null comment '消息标题',
    msg_type    int          default 0 comment '消息类型 1:info 2:success 3:warning 4:error',
    type        int          default 0 comment '消息类别 0:普通文本消息',
    text        varchar(500) not null comment '消息文本',
    param       json comment '消息上下文参数',
    template    varchar(500) default 'DefaultNotifyBody' comment '消息模板组件',
    status      int          default 0 comment '消息状态 0:未处理 1:已处理',
    is_delete   int          default 0 comment '是否删除',
    create_time timestamp(0) default now() comment '创建时间',
    update_time timestamp(0) default now() comment '处理时间'
) comment '消息通知表';





