BEGIN;
ALTER TABLE events
ADD created_by varchar(36) AFTER `deleted_at`;
COMMIT;