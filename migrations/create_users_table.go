package migrations

// CreateUsersTable is a migration that creates the users table.
// It implements the [migrate.Migration] interface.
type CreateUsersTable struct{}

// Up creates the users table.
func (CreateUsersTable) Up() string {
	return `
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

    CREATE TABLE IF NOT EXISTS public.users (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        username VARCHAR(255) UNIQUE NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        admin BOOLEAN DEFAULT FALSE,
        is_active BOOLEAN DEFAULT TRUE,
        is_deleted BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMP DEFAULT NOW(),
        updated_at TIMESTAMP DEFAULT NOW(),
        deleted_at TIMESTAMP
    );
    `
}

// Down drops the users table.
func (CreateUsersTable) Down() string {
	return `DROP TABLE IF EXISTS public.users;`
}

// Name returns the name of the migration.
func (CreateUsersTable) Name() string {
	return "CreateUsersTable"
}
