-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE admin (
    id SERIAL PRIMARY KEY,
    id_ktp VARCHAR(16) UNIQUE NOT NULL,
    nama_lengkap VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20),
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50) DEFAULT 'SYSTEM',
    modified_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(50) DEFAULT 'SYSTEM',
    active BOOLEAN DEFAULT true
);

CREATE TABLE dokter (
    id SERIAL PRIMARY KEY,
    id_ktp VARCHAR(16) NOT NULL,
    nama_lengkap VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(50),
    active BOOLEAN DEFAULT true
);

CREATE TABLE pasien (
    id SERIAL PRIMARY KEY,
    id_ktp VARCHAR(16) NOT NULL,
    no_rm VARCHAR(6) NOT NULL UNIQUE,
    nama_lengkap VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    tanggal_lahir DATE,
    golongan_darah VARCHAR(3),
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(50),
    active BOOLEAN DEFAULT true
);

CREATE TABLE obat (
    id SERIAL PRIMARY KEY,
    nama_obat VARCHAR(100) NOT NULL,
    jenis_obat VARCHAR(50) NOT NULL,
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(50),
    active BOOLEAN DEFAULT true
);

CREATE TABLE kunjungan (
    id SERIAL PRIMARY KEY,
    id_admin INTEGER REFERENCES admin(id),
    id_pasien INTEGER REFERENCES pasien(id),
    id_dokter INTEGER REFERENCES dokter(id),
    id_obat INTEGER REFERENCES obat(id),
    tanggal_kunjungan TIMESTAMP NOT NULL,
    riwayat_penyakit TEXT,
    diagnosa TEXT,
    resep_obat TEXT,
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(50),
    active BOOLEAN DEFAULT true
);

-- +migrate StatementEnd