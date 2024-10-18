-- +migrate Up
-- +migrate StatementBegin

-- Insert default admin user with bcrypt hashed password
INSERT INTO admin (
    id_ktp,
    nama_lengkap,
    email,
    password,
    role,
    created_time,
    created_by,
    modified_time,
    modified_by,
    active
)
VALUES (
    '0000000000000001',
    'Muhammad Ridwan Hakim, S.T., CPITA',
    'admin@rescenic.xyz',
    '$2a$10$kZJs0LrnmBhRhVR7zJSCEetNx25mxX/KY0pae7XNrNmxlRmWFFnva', -- Bcrypt hash for 'P@ssw0rd'
    'admin',
    NOW(),
    'SYSTEM',
    NOW(),
    'SYSTEM',
    true
);

-- +migrate StatementEnd
