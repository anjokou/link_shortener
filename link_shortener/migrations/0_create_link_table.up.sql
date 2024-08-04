CREATE TABLE links(
    id CHAR(8) PRIMARY KEY NOT NULL,
    external_url VARCHAR(2048) NOT NULL
);

CREATE UNIQUE INDEX idx_links_external_url ON links(external_url);
CREATE INDEX idx_links_id ON links(id);