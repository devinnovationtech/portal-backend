BEGIN;
ALTER TABLE news RENAME COLUMN author_id TO author;
ALTER TABLE news ADD COLUMN reporter VARCHAR(36) NULL AFTER author;
ALTER TABLE news ADD COLUMN editor VARCHAR (36) NULL AFTER reporter;
COMMIT;