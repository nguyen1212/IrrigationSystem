create table if not exists users(
    id text not null,
    uname text not null,
    email text not null,
    primary key (id)
);

create table if not exists presets(
    id text not null,
    mode text not null, 
    uid text not null,
    primary key (id),
    foreign key (uid) references users(id)
);

create table if not exists plots (
    id text not null,
    uid text not null,
    primary key (id),
    foreign key (uid) references users(id)
);

create table if not exists plot_presets(
    plot_id text not null,
    preset_id text not null,
    primary key (plot_id),
    foreign key (plot_id) references plots(id),
    foreign key (preset_id) references presets(id)
);

create table if not exists devices(
    id text not null,
    dname text not null,
    pid text not null,
    dtype text not null,
    primary key (id),
    foreign key (pid) references plots(id)
);

create table if not exists interval_presets(
    id text not null,
    amount decimal not null,
    prname text not null,
    interval decimal not null,
    primary key (id),
    foreign key (id) references presets(id)
);

create table if not exists repeat_presets(
    id text not null,
    prname text not null,
    water_time time not null,
    amount decimal not null,
    repeat text not null,
    primary key (id),
    foreign key (id) references presets(id)
);