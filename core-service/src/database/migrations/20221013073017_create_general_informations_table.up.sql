CREATE TABLE general_informations (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) DEFAULT NULL,
    alias VARCHAR(100) DEFAULT NULL,
    description TEXT DEFAULT NULL,
    slug VARCHAR(255) DEFAULT NULL,
    email VARCHAR(150) DEFAULT NULL,
    category VARCHAR(150) DEFAULT NULL,
    addresses json DEFAULT NULL,
    unit VARCHAR(150) DEFAULT NULL,
    phone json DEFAULT NULL,
    logo VARCHAR(255) DEFAULT NULL,
    operational_hours json DEFAULT NULL,
    media json DEFAULT NULL,
    social_media json DEFAULT NULL,
    link json DEFAULT NULL,
    type VARCHAR(100) DEFAULT NULL,    
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE INDEX idx_general_informations_name ON general_informations (name);
CREATE INDEX idx_general_informations_category ON general_informations (category);
CREATE INDEX idx_general_informations_type ON general_informations (type);