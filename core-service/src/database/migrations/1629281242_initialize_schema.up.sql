BEGIN;

DROP TABLE IF EXISTS areas;
CREATE TABLE areas (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  depth int(11) NULL DEFAULT NULL,
  name varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  parent_code_kemendagri varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  code_kemendagri varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  code_bps varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  latitude varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  longitude varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  meta varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (id) USING BTREE,
  UNIQUE INDEX areas_code_kemendagri_unique(code_kemendagri) USING BTREE,
  UNIQUE INDEX areas_code_bps_unique(code_bps) USING BTREE,
  INDEX areas_name_index(name) USING BTREE,
  INDEX areas_parent_code_kemendagri_index(parent_code_kemendagri) USING BTREE
);

-- ipj_db.units definition
DROP TABLE IF EXISTS units;
CREATE TABLE units (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  parent_id int(10) DEFAULT NULL,
  name varchar(100) NOT NULL,
  description varchar(255) DEFAULT NULL,
  logo varchar(255) DEFAULT NULL,
  website varchar(60) DEFAULT NULL,
  phone varchar(100) DEFAULT NULL,
  address varchar(255) DEFAULT NULL,
  chief varchar(100) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE INDEX idx_name ON units (name);

DROP TABLE IF EXISTS roles;
CREATE TABLE roles (
  id tinyint(2) UNSIGNED NOT NULL AUTO_INCREMENT,
  name varchar(100) NOT NULL,
  description varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id varchar(36) NOT NULL,
  name varchar(100) NOT NULL,
  username varchar(100) NOT NULL,
  email varchar(80) NOT NULL,
  password varchar(255) NOT NULL,
  unit_id int(10) unsigned,
  role_id tinyint(2) unsigned,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id),
  KEY users_unit_id_fk (unit_id),
  KEY users_role_id_fk (role_id),
  CONSTRAINT users_unit_id_fk FOREIGN KEY (unit_id) REFERENCES units (id),
  CONSTRAINT users_role_id_fk FOREIGN KEY (role_id) REFERENCES roles (id)
);
CREATE INDEX idx_username ON users(username);
CREATE INDEX idx_deleted_at ON users(deleted_at);

DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  title varchar(80) NOT NULL,
  description varchar(255),
  type varchar(80),
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS informations;
CREATE TABLE informations (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  category_id int(10) unsigned NOT NULL,
  title varchar(80) NOT NULL,
  excerpt varchar(150) NOT NULL,
  content text NOT NULL,
  slug varchar(100) DEFAULT NULL,
  image varchar(255) DEFAULT NULL,
  source varchar(80) DEFAULT NULL,
  show_date datetime NOT NULL,
  end_date datetime NOT NULL,
  status varchar(12) NOT NULL DEFAULT 'PUBLISHED',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY info_categories_id_fk (category_id),
  CONSTRAINT info_categories_id_fk FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_title ON informations (title);
CREATE INDEX idx_status ON informations (status);
CREATE INDEX idx_show_date ON informations (show_date);
CREATE INDEX idx_end_date ON informations (end_date);

DROP TABLE IF EXISTS news;
CREATE TABLE news (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  category varchar(30) NOT NULL,
  title varchar(80) NOT NULL,
  excerpt text NOT NULL,
  content text NOT NULL,
  slug varchar(100) UNIQUE NOT NULL,
  image varchar(255) DEFAULT NULL,
  video varchar(80) DEFAULT NULL,
  source varchar(80) DEFAULT NULL,
  status varchar(12) NOT NULL DEFAULT 'PUBLISHED',
  views bigint DEFAULT 0 NOT NULL,
  shared bigint DEFAULT 0 NOT NULL,
  highlight tinyint(1) NOT NULL,
  type varchar(20) NOT NULL DEFAULT 'article',
  tags varchar(100) NOT NULL,
  author_id varchar(36),
  created_by varchar(36),
  updated_by varchar(36),
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY users_author_id_fk (author_id),
  KEY users_created_by_fk (created_by),
  KEY users_updated_by_fk (updated_by),
  CONSTRAINT users_author_id_fk FOREIGN KEY (author_id) REFERENCES users (id),
  CONSTRAINT users_created_by_fk FOREIGN KEY (created_by) REFERENCES users (id),
  CONSTRAINT users_updated_by_fk FOREIGN KEY (updated_by) REFERENCES users (id)
);
CREATE INDEX idx_title ON news (title);
CREATE INDEX idx_category ON news (category);
CREATE INDEX idx_status ON news (status);
CREATE INDEX news_views_index ON news (views);

DROP TABLE IF EXISTS events;
CREATE TABLE events (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  title varchar(80) NOT NULL,
  description varchar(255),
  date date NOT NULL,
  priority tinyint(1) DEFAULT 1 NOT NULL,
  start_hour time NOT NULL,
  end_hour time NOT NULL,
  image varchar(255) DEFAULT NULL,
  published_by varchar(16) DEFAULT NULL,
  type ENUM('offline', 'online') NOT NULL,
  address varchar(255) DEFAULT NULL,
  url varchar(80) DEFAULT NULL,
  category varchar(30) NOT NULL,
  province_code varchar(191) NULL,
  city_code varchar(191) NULL,
  district_code varchar(191) NULL,
  village_code varchar(191) NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
);
CREATE INDEX idx_title ON events (title);
CREATE INDEX idx_start_hour ON events (start_hour);
CREATE INDEX idx_end_hour ON events (end_hour);
CREATE INDEX idx_category ON events (category);

DROP TABLE IF EXISTS feedback;
CREATE TABLE feedback (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  rating int(11) DEFAULT NULL,
  compliments text NOT NULL,
  criticism text NOT NULL,
  suggestions text NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS featured_programs;
CREATE TABLE featured_programs (
	id int(10) unsigned NOT NULL AUTO_INCREMENT,
	title varchar(100) not null,
	excerpt varchar(255) not null,
	description text not null,
	organization varchar(100),
	categories json,
	service_type varchar(10),
	websites json,
	social_media json,
	logo varchar(150),
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE INDEX idx_categories_program on featured_programs (categories);

DROP TABLE IF EXISTS tags;
CREATE TABLE tags (
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(20) NOT NULL,
    PRIMARY KEY(id)
);
CREATE INDEX idx_tags_name ON tags (name);

DROP TABLE IF EXISTS tags_data;
CREATE TABLE tags_data (
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    data_id int(10) unsigned,
    tags_id int(10) unsigned,
    tags_name varchar(20),
    type varchar(10),
    PRIMARY KEY(id)
);
CREATE INDEX idx_tags_name ON tags_data (tags_name);

COMMIT;
