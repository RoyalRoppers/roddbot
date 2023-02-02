CREATE TABLE guild_role_mappings (
    guild_id TEXT NOT NULL PRIMARY KEY REFERENCES guilds(id),
    admin_role_id TEXT,
    player_role_id TEXT
);
