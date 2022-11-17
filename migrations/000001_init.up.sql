CREATE TABLE guilds (
    id TEXT NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE ctf_channels (
    id TEXT PRIMARY KEY,
    topic_chan TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    guild_id TEXT REFERENCES guilds(id) NOT NULL,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    archived BOOLEAN NOT NULL,
    username TEXT,
    password TEXT,
    ctftime_id INT
);

CREATE TABLE chall_channels (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    parent_id TEXT NOT NULL REFERENCES ctf_channels(id),
    title TEXT NOT NULL,
    flag TEXT,
    solved_at TIMESTAMPTZ
);