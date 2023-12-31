BEGIN;
CREATE TABLE registration_invitations (
  id varchar(36) NOT NULL,
  email VARCHAR(80) NOT NULL,
  token VARCHAR(255),
  unit_id int(10),
  invited_by varchar(36) NOT NULL,
  invited_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE INDEX idx_email ON registration_invitations(email);
CREATE INDEX idx_token ON registration_invitations(token);
COMMIT;