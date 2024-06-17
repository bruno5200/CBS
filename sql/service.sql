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

INSERT INTO storage.services (
service_id,
service_name,
service_key,
service_description
) VALUES
('cd05d13d-6555-42af-ae1e-dce46884d807','pagos','$2a$04$d9VJdySAxrv6O6j.P74Gju.OEYRHK0yCeO5JD/rifAIp84JG7dABq','Almacenamiento de archivos Pasarela de Pagos - Datec');