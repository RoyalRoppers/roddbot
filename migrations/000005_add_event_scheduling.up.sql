CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title TEXT NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    -- discord events can have voice channels connected to them instead of a location??
    location TEXT NOT NULL,
    description TEXT NOT NULL,
    announcement_channel TEXT NOT NULL,
    announcement_time TIMESTAMPTZ,
    announcement_msg_id TEXT,
    announcement_event_id TEXT
);
-- https://github.com/arran4/golang-ical
