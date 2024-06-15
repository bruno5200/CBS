-- Create table storage.blocks
CREATE TABLE IF NOT EXISTS storage.blocks (
	block_id UUID NOT NULL,
	block_name TEXT NOT NULL,
	block_checksum TEXT NOT NULL,
	block_extension TEXT NOT NULL,
	block_url TEXT NOT NULL,
	block_uploaded_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	block_group_id UUID NOT NULL,
	block_service_id UUID NOT NULL,
	state BOOLEAN NOT NULL DEFAULT TRUE,
	CONSTRAINT block_id_pk PRIMARY KEY (block_id),
	CONSTRAINT block_group_id_fk FOREIGN KEY (block_group_id) REFERENCES storage.groups (group_id),
	CONSTRAINT block_service_id_fk FOREIGN KEY (block_service_id) REFERENCES storage.services (service_id)
);

-- Functions for blocks management

-- Create block
CREATE OR REPLACE FUNCTION storage.create_block(
	_id
	_name TEXT,
	_checksum TEXT,
	_extension TEXT,
	_url TEXT,
	_group_id UUID,
	_service_id UUID
) RETURNS VOID
AS
$BODY$
BEGIN
	INSERT INTO (
		block_id,
		block_name,
		block_checksum,
		block_extension,
		block_url,
		block_group_id,
		block_service_id
	) VALUES (
		_id,
		_name,
		_checksum,
		_extension,
		_url,
		_group_id,
		_service_id
	);
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Block not created';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;

-- Read block
CREATE OR REPLACE FUNCTION storage.read_block(
	_id UUID
)
RETURNS TABLE (
	_id UUID,
	_name TEXT,
	_checksum TEXT,
	_extension TEXT,
	_url TEXT,
	_group_id UUID,
	_service_id UUID
)
AS
$BODY$
BEGIN
	RETURN QUERY SELECT
		block_id,
		block_name,
		block_checksum,
		block_extension,
		block_url,
		block_group_id,
		block_service_id
	FROM storage.blocks
	WHERE block_id = _id;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Block not found';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;

-- Read Blocks by Group
CREATE OR REPLACE FUNCTION storage.read_blocks_by_group(
	_group_id UUID
)
RETURNS TABLE (
	_id UUID,
	_name TEXT,
	_checksum TEXT,
	_extension TEXT,
	_url TEXT,
	_group_id UUID,
	_service_id UUID
)
AS
$BODY$
BEGIN
	RETURN QUERY SELECT
		block_id,
		block_name,
		block_checksum,
		block_extension,
		block_url,
		block_group_id,
		block_service_id
	FROM storage.blocks
	WHERE block_group_id = _group_id;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Blocks not found';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;

-- Read Blocks by Service
CREATE OR REPLACE FUNCTION storage.read_blocks_by_service(
	_service_id UUID
)
RETURNS TABLE (
	_id UUID,
	_name TEXT,
	_checksum TEXT,
	_extension TEXT,
	_url TEXT,
	_group_id UUID,
	_service_id UUID
)
AS
$BODY$
BEGIN
	RETURN QUERY SELECT
		block_id,
		block_name,
		block_checksum,
		block_extension,
		block_url,
		block_group_id,
		block_service_id
	FROM storage.blocks
	WHERE block_service_id = _service_id;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Blocks not found';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;

-- Update block
CREATE OR REPLACE FUNCTION storage.update_block(
	_id UUID,
	_name TEXT,
	_checksum TEXT,
	_extension TEXT,
	_url TEXT,
	_group_id UUID,
	_service_id UUID
)
RETURNS VOID
AS
$BODY$
BEGIN
	UPDATE storage.blocks
	SET
		block_name = _name,
		block_checksum = _checksum,
		block_extension = _extension,
		block_url = _url,
		block_group_id = _group_id,
		block_service_id = _service_id
	WHERE block_id = _id;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Block not updated';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;

-- Disable block
CREATE OR REPLACE FUNCTION storage.disable_block(
	_id UUID
)
RETURNS VOID
AS
$BODY$
BEGIN
	UPDATE storage.blocks
	SET state = FALSE
	WHERE block_id = _id;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Block not disabled';
	END IF;
END;
