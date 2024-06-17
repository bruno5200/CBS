-- Create table storage.groups
CREATE TABLE IF NOT EXISTS storage.groups (
    group_id UUID NOT NULL,
    group_name TEXT NOT NULL,
    group_description TEXT NOT NULL,
    group_service_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    delete_at TIMESTAMPTZ,
    state BOOLEAN NOT NULL DEFAULT TRUE,
    CONSTRAINT group_id_pk PRIMARY KEY (group_id),
    CONSTRAINT group_service_id_fk FOREIGN KEY (group_service_id) REFERENCES storage.services (service_id)
);

INSERT INTO storage.groups (
    group_id,
    group_name,
    group_description
) VALUES
('8a66950b-fbcb-4f9b-9361-86e8392e043f','Test','','cd05d13d-6555-42af-ae1e-dce46884d807');

-- Functions for groups management

-- Create group
CREATE OR REPLACE FUNCTION storage.fn_create_group(
    _id UUID,
    _name TEXT,
    _description TEXT,
    _service_id UUID
)
RETURNS VOID
AS
$BODY$
BEGIN
    INSERT INTO storage.groups(
        group_id,
        group_name,
        group_description,
        group_service_id
    ) VALUES (
        _id,
        _name,
        _description,
        _service_id
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
    description TEXT,
    service_id UUID,
    service_name TEXT
)
AS
$BODY$
BEGIN
    RETURN QUERY
    SELECT
        g.group_id,
        g.group_name,
        g.group_description,
    FROM storage.groups AS g
    INNER JOIN storage.services AS s ON g.group_service_id = s.service_id
    WHERE group_id = _id;
    IF NOT FOUND THEN
        RAISE EXCEPTION 'Group not found';
    END IF;
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
    WHERE b.group_service_id = _service_id
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