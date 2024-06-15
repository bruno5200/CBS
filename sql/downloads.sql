-- Create table storage.downloads
CREATE TABLE IF NOT EXISTS storage.downloads (
    download_id UUID NOT NULL,
    downloaded_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT download_id_pk PRIMARY KEY (download_id),
    CONSTRAINT download_id_fk FOREIGN KEY (download_id) REFERENCES storage.blocks (block_id)
);