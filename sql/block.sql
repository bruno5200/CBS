-- Create table storage.blocks
CREATE TABLE IF NOT EXISTS storage.blocks (
	block_id UUID NOT NULL,
	block_name TEXT NOT NULL,
	block_checksum TEXT NOT NULL,
	block_extension TEXT NOT NULL,
	block_url TEXT NOT NULL,
	block_uploaded_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	block_group_id UUID NOT NULL,
	state BOOLEAN NOT NULL DEFAULT TRUE,
	CONSTRAINT block_id_pk PRIMARY KEY (block_id),
	CONSTRAINT block_group_id_fk FOREIGN KEY (block_group_id) REFERENCES storage.groups (group_id)
);

INSERT INTO storage.blocks (
	block_id,
	block_name,
	block_checksum,
	block_extension,
	block_url,
	block_group_id
) VALUES
('9add6cca-25f8-4657-8e88-0bf7f9a12cbb','reporte_solicitudes_01_12_2023_31_12_2023.csv','3d7907504faa776990fa9c01e3ff2dd1ae391890b972bca664366b59cbc5b53c','CSV','https://blob.gutier.lat/documents/9add6cca-25f8-4657-8e88-0bf7f9a12cbb.csv','8a66950b-fbcb-4f9b-9361-86e8392e043f'),
('10893629-aa43-4552-bd22-870bc85a5bea','CARLOS DANIEL VILLALBA RADA_ACTA DE ENTREGA_FIRMADO.pdf','ea971396ea93d5515b6cef3307115e0f540dc9dee24ca5ca6fd297d3635072ca','PDF','https://blob.gutier.lat/documents/10893629-aa43-4552-bd22-870bc85a5bea.pdf','8a66950b-fbcb-4f9b-9361-86e8392e043f');

-- Functions for blocks management

-- Create block
CREATE OR REPLACE FUNCTION storage.fn_create_block(
	_id UUID,
	_name TEXT,
	_checksum TEXT,
	_extension TEXT,
	_url TEXT,
	_group_id UUID
) RETURNS VOID
AS
$BODY$
BEGIN
	INSERT INTO storage.blocks(
		block_id,
		block_name,
		block_checksum,
		block_extension,
		block_url,
		block_group_id
	) VALUES (
		_id,
		_name,
		_checksum,
		_extension,
		_url,
		_group_id
	);
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Block not created';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;
-- SELECT storage.fn_create_block('9add6cca-25f8-4657-8e88-0bf7f9a12cbb','reporte_solicitudes_01_12_2023_31_12_2023.csv','3d7907504faa776990fa9c01e3ff2dd1ae391890b972bca664366b59cbc5b53c','CSV','https://blob.gutier.lat/documents/9add6cca-25f8-4657-8e88-0bf7f9a12cbb.csv','8a66950b-fbcb-4f9b-9361-86e8392e043f');

-- Read block
CREATE OR REPLACE FUNCTION storage.fn_read_block(
	_id UUID
)
RETURNS TABLE (
	id UUID,
	name TEXT,
	checksum TEXT,
	extension TEXT,
	url TEXT,
	at TIMESTAMPTZ,
	group_id UUID,
	group_name TEXT,
	service_id UUID,
	service_name TEXT,
	active BOOLEAN
)
AS
$BODY$
BEGIN
	RETURN QUERY SELECT
	b.block_id,
	b.block_name,
	b.block_checksum,
	b.block_extension,
	b.block_url,
	b.block_uploaded_at,
	g.group_id,
	g.group_name,
	s.service_id,
	s.service_name,
	b.state
	FROM storage.blocks AS b
	INNER JOIN storage.groups g ON b.block_group_id = g.group_id
	INNER JOIN storage.services s ON g.group_service_id = s.service_id
	WHERE b.block_id = _id
	AND g.state;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Block not found';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;
-- SELECT id, name, checksum, extension, url, at, group_id, group_name, service_id, service_name, active FROM storage.fn_read_block('9add6cca-25f8-4657-8e88-0bf7f9a12cbb');

-- Read Block by Checksum
CREATE OR REPLACE FUNCTION storage.fn_read_block_by_checksum(
	_checksum TEXT
)
RETURNS TABLE (
	id UUID,
	name TEXT,
	checksum TEXT,
	extension TEXT,
	url TEXT,
	at TIMESTAMPTZ,
	group_id UUID,
	group_name TEXT,
	service_id UUID,
	service_name TEXT,
	active BOOLEAN
)
AS
$BODY$
BEGIN
	RETURN QUERY SELECT
	b.block_id,
	b.block_name,
	b.block_checksum,
	b.block_extension,
	b.block_url,
	b.block_uploaded_at,
	g.group_id,
	g.group_name,
	s.service_id,
	s.service_name,
	b.state
	FROM storage.blocks AS b
	INNER JOIN storage.groups g ON b.block_group_id = g.group_id
	INNER JOIN storage.services s ON g.group_service_id = s.service_id
	WHERE b.block_checksum = _checksum
	AND g.state;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Block not found';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;
-- SELECT id, name, checksum, extension, url, at, group_id, group_name, service_id, service_name, active FROM storage.fn_read_block_by_checksum('3d7907504faa776990fa9c01e3ff2dd1ae391890b972bca664366b59cbc5b53c');

-- Read Blocks by Group
CREATE OR REPLACE FUNCTION storage.fn_read_blocks_by_group(
	_group_id UUID
)
RETURNS TABLE (
	id UUID,
	name TEXT,
	checksum TEXT,
	extension TEXT,
	url TEXT,
	at TIMESTAMPTZ,
	group_id UUID,
	group_name TEXT,
	service_id UUID,
	service_name TEXT,
	active BOOLEAN
)
AS
$BODY$
BEGIN
	RETURN QUERY SELECT
	b.block_id,
	b.block_name,
	b.block_checksum,
	b.block_extension,
	b.block_url,
	b.block_uploaded_at,
	g.group_id,
	g.group_name,
	s.service_id,
	s.service_name,
	b.state
	FROM storage.blocks AS b
	INNER JOIN storage.groups g ON b.block_group_id = g.group_id
	INNER JOIN storage.services s ON g.group_service_id = s.service_id
	WHERE g.group_id = _group_id
	AND g.state;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Blocks not found';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;
-- SELECT id, name, checksum, extension, url, at, group_id, group_name, service_id, service_name, active FROM storage.fn_read_blocks_by_group('8a66950b-fbcb-4f9b-9361-86e8392e043f');

-- Read Blocks by Service
CREATE OR REPLACE FUNCTION storage.fn_read_blocks_by_service(
	_service_id UUID
)
RETURNS TABLE (
	id UUID,
	name TEXT,
	checksum TEXT,
	extension TEXT,
	url TEXT,
	at TIMESTAMPTZ,
	group_id UUID,
	group_name TEXT,
	service_id UUID,
	service_name TEXT,
	active BOOLEAN
)
AS
$BODY$
BEGIN
	RETURN QUERY SELECT
	b.block_id,
	b.block_name,
	b.block_checksum,
	b.block_extension,
	b.block_url,
	b.block_uploaded_at,
	g.group_id,
	g.group_name,
	s.service_id,
	s.service_name,
	b.state
	FROM storage.blocks AS b
	INNER JOIN storage.groups g ON b.block_group_id = g.group_id
	INNER JOIN storage.services s ON g.group_service_id = s.service_id
	WHERE s.service_id = _service_id
	AND g.state;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Blocks not found';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;
-- SELECT id, name, checksum, extension, url, at, group_id, group_name, service_id, service_name, active FROM storage.fn_read_blocks_by_service('cd05d13d-6555-42af-ae1e-dce46884d807');

-- Update block
CREATE OR REPLACE FUNCTION storage.fn_update_block(
	_id UUID,
	_name TEXT,
	_checksum TEXT,
	_extension TEXT,
	_url TEXT,
	_group_id UUID,
	_active BOOLEAN
)
RETURNS VOID
AS
$BODY$
BEGIN
	UPDATE storage.blocks AS b
	SET
		block_name = _name,
		block_checksum = _checksum,
		block_extension = _extension,
		block_url = _url,
		block_group_id = _group_id,
		state = _active
	WHERE b.block_id = _id
	AND b.state;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Block not updated';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;
-- SELECT storage.fn_update_block('9add6cca-25f8-4657-8e88-0bf7f9a12cbb','reporte_solicitudes_01_12_2023_31_12_2023.csv','3d7907504faa776990fa9c01e3ff2dd1ae391890b972bca664366b59cbc5b53c','CSV','https://blob.gutier.lat/documents/9add6cca-25f8-4657-8e88-0bf7f9a12cbb.csv','8a66950b-fbcb-4f9b-9361-86e8392e043f',TRUE);

-- Disable block
CREATE OR REPLACE FUNCTION storage.fn_disable_block(
	_id UUID
)
RETURNS VOID
AS
$BODY$
BEGIN
	UPDATE storage.blocks AS b SET
	state = FALSE
	WHERE b.block_id = _id
	AND b.state;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Block not disabled';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;
-- SELECT storage.fn_disable_block('9add6cca-25f8-4657-8e88-0bf7f9a12cbb');