-- Create table storage.downloads
CREATE TABLE IF NOT EXISTS storage.downloads (
    download_id UUID NOT NULL,
    downloaded_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT download_id_pk PRIMARY KEY (download_id),
    CONSTRAINT download_id_fk FOREIGN KEY (download_id) REFERENCES storage.blocks (block_id)
);

-- Functions for downloads management

-- Create download
CREATE OR REPLACE FUNCTION storage.fn_create_download(
    _id UUID
) RETURNS VOID
AS
$BODY$
BEGIN
    INSERT INTO storage.downloads(
        download_id
    ) VALUES (
        _id
    );
    IF NOT FOUND THEN
        RAISE EXCEPTION 'Download not created';
    END IF;
END;
$BODY$
LANGUAGE plpgsql;