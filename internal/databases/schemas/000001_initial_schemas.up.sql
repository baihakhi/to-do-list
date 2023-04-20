CREATE TABLE IF NOT EXISTS users (
    id bigserial NOT NULL,
    code varchar(100) NOT NULL UNIQUE,
    name varchar(255) DEFAULT '',
    user_group varchar(50) DEFAULT 'MEMBER' NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    token varchar(255),
    status varchar(100) DEFAULT 'INACTIVE' NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS activities (
    id bigserial NOT NULL,
    code varchar(100) NOT NULL UNIQUE,
    name varchar(255) NOT NULL,
    description text DEFAULT '',
    owner bigserial NOT NULL,
    assignee bigserial,
    is_deleted boolean DEFAULT false NOT NULL,
    deleted_at TIMESTAMP,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT activities_pkey PRIMARY KEY (id),
    FOREIGN KEY (assignee) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (owner) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE  FUNCTION update_timestamp_func()
RETURNS TRIGGER
LANGUAGE plpgsql AS
'
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
';

CREATE  FUNCTION create_timestamp_func()
RETURNS TRIGGER
LANGUAGE plpgsql AS
'
BEGIN
    NEW.created_at = now();
    NEW.updated_at = now();
    RETURN NEW;
END;
';

DO $$
DECLARE
    t text;
BEGIN
    FOR t IN
        SELECT table_name FROM information_schema.columns WHERE column_name = 'updated_at'
    LOOP
        EXECUTE format('CREATE TRIGGER trigger_update_timestamp
                    BEFORE UPDATE ON %I
                    FOR EACH ROW EXECUTE PROCEDURE update_timestamp_func()', t,t);
    END loop;

    FOR t IN
        SELECT table_name FROM information_schema.columns WHERE column_name = 'created_at'
    LOOP
        EXECUTE format('CREATE TRIGGER trigger_create_timestamp
                    BEFORE INSERT ON %I
                    FOR EACH ROW EXECUTE PROCEDURE create_timestamp_func()', t,t);
    END loop;
END;
$$ language 'plpgsql';
