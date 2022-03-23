

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    username VARCHAR(150) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    picture VARCHAR(256) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS employee_type (
    id BIGSERIAL NOT NULL,
    name VARCHAR(150) NOT NULL UNIQUE,
    description VARCHAR(150),
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_employee_type PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS workplace (
    id BIGSERIAL NOT NULL,
    name VARCHAR(150) NOT NULL,
    code VARCHAR(150) NOT NULL UNIQUE,
    address VARCHAR(150),
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_workplace PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS employee (
    id BIGSERIAL NOT NULL,
    file_code VARCHAR(150) UNIQUE,
    agent_number VARCHAR(150) UNIQUE,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) DEFAULT '',
    document_number VARCHAR(150) UNIQUE,
    birth_date DATE,
    date_admission DATE,
    phone VARCHAR(150) DEFAULT '',
    address VARCHAR(150) DEFAULT '',
    picture VARCHAR(256) DEFAULT '',
    salary NUMERIC(10,2) DEFAULT 0,
    category INTEGER DEFAULT 0,
    status INTEGER DEFAULT 1,
    work_number VARCHAR(150) DEFAULT '',
    employee_type BIGINT,
    workplace BIGINT,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    CONSTRAINT pk_employee PRIMARY KEY(id),
    CONSTRAINT fk_employee_type FOREIGN KEY(employee_type) 
    REFERENCES employee_type(id) ON DELETE CASCADE,
    CONSTRAINT fk_workplace FOREIGN KEY(workplace) 
    REFERENCES workplace(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS paying (
    id BIGSERIAL NOT NULL,
    employee_id BIGINT NOT NULL,
    month VARCHAR(150) NOT NULL,
    year VARCHAR(150) NOT NULL,
    amount NUMERIC(10,2) NOT NULL,
    description VARCHAR(350) NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    CONSTRAINT pk_paying PRIMARY KEY(id),
    CONSTRAINT fk_employee FOREIGN KEY(employee_id)
    REFERENCES employee(id)
);