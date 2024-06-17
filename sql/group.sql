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

INSERT INTO storage.groups (
    group_id,
    group_name,
    group_description
) VALUES
('8a66950b-fbcb-4f9b-9361-86e8392e043f','Test','');

-- Functions for groups management

-- Create group
CREATE OR REPLACE FUNCTION storage.fn_create_group(
    _id UUID,
    _name TEXT,
    _description TEXT
)
RETURNS VOID
AS
$BODY$
BEGIN
    INSERT INTO storage.groups(
        group_id,
        group_name,
        group_description
    ) VALUES (
        _id,
        _name,
        _description
    );
    IF NOT FOUND THEN
        RAISE EXCEPTION 'Group not created';
    END IF;
END;
$BODY$
LANGUAGE plpgsql;
-- SELECT storage.fn_create_group('8a66950b-fbcb-4f9b-9361-86e8392e043f','Test','');

-- Read group
CREATE OR REPLACE FUNCTION storage.fn_read_group(
    _id UUID
)
RETURNS TABLE (
    id UUID,
    name TEXT,
    description TEXT
)
AS
$BODY$
BEGIN
    RETURN QUERY
    SELECT
        group_id,
        group_name,
        group_description
    FROM storage.groups
    WHERE group_id = _id;
END;
$BODY$
LANGUAGE plpgsql;
-- SELECT id, name, description FROM storage.fn_read_group('8a66950b-fbcb-4f9b-9361-86e8392e043f');

-- Read groups by service
CREATE OR REPLACE FUNCTION storage.fn_read_groups_by_service(
    _service_id UUID
)
RETURNS TABLE (
    id UUID,
    name TEXT,
    description TEXT
)
AS
$BODY$
BEGIN
    RETURN QUERY
    SELECT
        g.group_id,
        g.group_name,
        g.group_description
    FROM storage.groups g
    INNER JOIN storage.blocks b ON g.group_id = b.block_group_id
    WHERE b.block_service_id = _service_id
    AND g.state
    AND b.state
    GROUP BY g.group_id;
END;
$BODY$
LANGUAGE plpgsql;
-- SELECT id, name, description FROM storage.fn_read_groups_by_service('cd05d13d-6555-42af-ae1e-dce46884d807');

-- Update group
CREATE OR REPLACE FUNCTION storage.fn_update_group(
    _id UUID,
    _name TEXT,
    _description TEXT
)
RETURNS VOID
AS
$BODY$
BEGIN
    UPDATE storage.groups
    SET
        group_name = _name,
        group_description = _description,
        updated_at = NOW()
    WHERE group_id = _id;
    IF NOT FOUND THEN
        RAISE EXCEPTION 'Group not updated';
    END IF;
END;
$BODY$
LANGUAGE plpgsql;
-- SELECT storage.fn_update_group('8a66950b-fbcb-4f9b-9361-86e8392e043f','Test','');

-- Disable group
CREATE OR REPLACE FUNCTION storage.fn_disable_group(
    _id UUID
)
RETURNS VOID
AS
$BODY$
BEGIN
    UPDATE storage.groups SET
        state = FALSE,
        delete_at = NOW()
    WHERE group_id = _id;
    IF NOT FOUND THEN
        RAISE EXCEPTION 'Group not disabled';
    END IF;
END;
$BODY$
LANGUAGE plpgsql;