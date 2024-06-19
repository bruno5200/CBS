-- Create table storage.services
CREATE TABLE IF NOT EXISTS storage.services (
	service_id UUID NOT NULL,
	service_name TEXT NOT NULL,
	service_key TEXT NOT NULL,
	service_description TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ,
	delete_at TIMESTAMPTZ,
	state BOOLEAN NOT NULL DEFAULT TRUE,
	CONSTRAINT service_id_pk PRIMARY KEY (service_id),
	CONSTRAINT service_name_uk UNIQUE (service_name),
	CONSTRAINT service_key_uk UNIQUE (service_key)
	CONSTRAINT service_name_lowercase_ck CHECK (service_name = LOWER(service_name))
);

INSERT INTO storage.services (
service_id,
service_name,
service_key,
service_description
) VALUES
('cd05d13d-6555-42af-ae1e-dce46884d807','pagos','$2a$04$d9VJdySAxrv6O6j.P74Gju.OEYRHK0yCeO5JD/rifAIp84JG7dABq','Almacenamiento de archivos Pasarela de Pagos - Datec');

-- Functions for service management

-- Create service
CREATE OR REPLACE FUNCTION storage.create_service(
	_id
	_name TEXT,
	_key TEXT,
	_description TEXT
)
RETURNS VOID
AS
$BODY$
BEGIN
	INSERT INTO storage.services (
		service_id,
		service_name,
		service_key,
		service_description
	) VALUES (
		_id,
		_name,
		_key,
		_description
	);
END;
$BODY$
LANGUAGE plpgsql;

-- Read service
CREATE OR REPLACE FUNCTION storage.read_service(
	_id UUID
)
RETURNS TABLE (
	service_id UUID,
	service_name TEXT,
	service_key TEXT,
	service_description TEXT
	state BOOLEAN
)
AS
$BODY$
BEGIN
	RETURN QUERY
	SELECT
		service_id,
		service_name,
		service_key,
		service_description,
		state
	FROM storage.services
	WHERE service_id = _id;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Service not found';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;

-- Read service by key
CREATE OR REPLACE FUNCTION storage.read_service_by_key(
	_key TEXT
)
RETURNS TABLE (
	service_id UUID,
	service_name TEXT,
	service_key TEXT,
	service_description TEXT
	state BOOLEAN
)
AS
$BODY$
BEGIN
	RETURN QUERY
	SELECT
		service_id,
		service_name,
		service_key,
		service_description,
		state
	FROM storage.services
	WHERE service_key = _key;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Service not found';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;

-- Update service
CREATE OR REPLACE FUNCTION storage.update_service(
	_id UUID,
	_name TEXT,
	_key TEXT,
	_description TEXT
)
RETURNS VOID
AS
$BODY$
BEGIN
	UPDATE storage.services
	SET
		service_name = _name,
		service_key = _key,
		service_description = _description,
		updated_at = NOW()
	WHERE service_id = _id;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Service not found';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;

-- Disable service
CREATE OR REPLACE FUNCTION storage.disable_service(
	_id UUID
)
RETURNS VOID
AS
$BODY$
BEGIN
	UPDATE storage.services
	SET
		state = FALSE,
		updated_at = NOW()
	WHERE service_id = _id;
	IF NOT FOUND THEN
		RAISE EXCEPTION 'Service not found';
	END IF;
END;
$BODY$
LANGUAGE plpgsql;