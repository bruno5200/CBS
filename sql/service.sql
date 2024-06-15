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
	CONSTRAINT service_key_unique UNIQUE (service_key)
);