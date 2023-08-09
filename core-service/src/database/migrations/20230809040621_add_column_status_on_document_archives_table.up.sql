ALTER TABLE document_archives
ADD COLUMN status varchar(255) DEFAULT "PUBLISHED" AFTER `category`;
