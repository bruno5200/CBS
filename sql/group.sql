-- Create table storage.groups
CREATE TABLE IF NOT EXISTS storage.groups (
    group_id UUID NOT NULL,
    group_name TEXT NOT NULL,
    group_description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    delete_at TIMESTAMPTZ,
    state BOOLEAN NOT NULL DEFAULT TRUE,
    CONSTRAINT group_id_pk PRIMARY KEY (group_id)
);