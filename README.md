# Gumuruh Clinic API Guide

## Base URL

```sh
http://localhost:8090/api
```

## Authentication

- Use **JWT** for authentication.
- Obtain a token by logging in with the `/login` endpoint.

## API Endpoints

### Public Routes

#### 1. Register a User

- **Endpoint**: `POST /register`
- **Request Body**:

  ```json
  {
    "id_ktp": "1234567890123456",
    "nama_lengkap": "John Doe",
    "email": "johndoe@example.com",
    "password": "yourpassword",
    "tanggal_lahir": "2000-01-01",
    "golongan_darah": "O"
  }
  ```

- **Response**:
  - **201 Created**: User successfully registered.
  - **400 Bad Request**: Validation errors.

#### 2. Login

- **Endpoint**: `POST /login`
- **Request Body**:

  ```json
  {
    "email": "johndoe@example.com",
    "password": "yourpassword",
    "role": "pasien" // or "admin", "dokter"
  }
  ```

- **Response**:
  - **200 OK**: Returns a JWT token.
  - **401 Unauthorized**: Invalid credentials.

### Protected Routes (Requires JWT Token)

#### 3. Admin Routes

##### Create Admin

- **Endpoint**: `POST /admin`
- **Headers**:

  ```http
  Authorization: Bearer <your_jwt_token>
  ```

- **Request Body**:

  ```json
  {
    "id_ktp": "1234567890123456",
    "nama_lengkap": "Admin User",
    "email": "admin@example.com",
    "password": "adminpassword"
  }
  ```

- **Response**:
  - **201 Created**: Admin successfully created.

##### List Admin

- **Endpoint**: `GET /admin`
- **Response**:
  - **200 OK**: Returns a list of admins.

##### Get Admin by ID

- **Endpoint**: `GET /admin/:id`
- **Response**:
  - **200 OK**: Returns admin details.
  - **404 Not Found**: Admin not found.

##### Update Admin

- **Endpoint**: `PUT /admin/:id`
- **Request Body**: Same as Create Admin
- **Response**:
  - **200 OK**: Admin successfully updated.

##### Delete Admin

- **Endpoint**: `DELETE /admin/:id`
- **Response**:
  - **200 OK**: Admin successfully deleted.

#### 4. Pasien Routes

##### Create Pasien

- **Endpoint**: `POST /pasien`
- **Headers**:

  ```http
  Authorization: Bearer <your_jwt_token>
  ```

- **Request Body**: Same as Register
- **Response**:
  - **201 Created**: Pasien successfully created.

##### List Pasien

- **Endpoint**: `GET /pasien`
- **Response**:
  - **200 OK**: Returns a list of Pasiens.

##### Get Pasien by ID

- **Endpoint**: `GET /pasien/:id`
- **Response**:
  - **200 OK**: Returns pasien details.

##### Update Pasien

- **Endpoint**: `PUT /pasien/:id`
- **Request Body**: Same as Create Pasien
- **Response**:
  - **200 OK**: Pasien successfully updated.

##### Delete Pasien

- **Endpoint**: `DELETE /pasien/:id`
- **Response**:
  - **200 OK**: Pasien successfully deleted.

#### 5. Dokter Routes

##### Create Dokter

- **Endpoint**: `POST /dokter`
- **Headers**:

  ```http
  Authorization: Bearer <your_jwt_token>
  ```

- **Request Body**: Same as Register
- **Response**:
  - **201 Created**: Dokter successfully created.

##### List Dokter

- **Endpoint**: `GET /dokter`
- **Response**:
  - **200 OK**: Returns a list of Dokters.

##### Get Dokter by ID

- **Endpoint**: `GET /dokter/:id`
- **Response**:
  - **200 OK**: Returns dokter details.

##### Update Dokter

- **Endpoint**: `PUT /dokter/:id`
- **Request Body**: Same as Create Dokter
- **Response**:
  - **200 OK**: Dokter successfully updated.

##### Delete Dokter

- **Endpoint**: `DELETE /dokter/:id`
- **Response**:
  - **200 OK**: Dokter successfully deleted.

#### 6. Obat Routes

##### Create Obat

- **Endpoint**: `POST /obat`
- **Headers**:

  ```http
  Authorization: Bearer <your_jwt_token>
  ```

- **Request Body**: Same as Register
- **Response**:
  - **201 Created**: Obat successfully created.

##### List Obat

- **Endpoint**: `GET /obat`
- **Response**:
  - **200 OK**: Returns a list of Obats.

##### Get Obat by ID

- **Endpoint**: `GET /obat/:id`
- **Response**:
  - **200 OK**: Returns obat details.

##### Update Obat

- **Endpoint**: `PUT /obat/:id`
- **Request Body**: Same as Create Obat
- **Response**:
  - **200 OK**: Obat successfully updated.

##### Delete Obat

- **Endpoint**: `DELETE /obat/:id`
- **Response**:
  - **200 OK**: Obat successfully deleted.

#### 7. Kunjungan Routes

##### Create Kunjungan

- **Endpoint**: `POST /kunjungan`
- **Headers**:

  ```http
  Authorization: Bearer <your_jwt_token>
  ```

- **Request Body**: Same as Register
- **Response**:
  - **201 Created**: Kunjungan successfully created.

##### List Kunjungan

- **Endpoint**: `GET /kunjungan`
- **Response**:
  - **200 OK**: Returns a list of Kunjungans.

##### Get Kunjungan by ID

- **Endpoint**: `GET /kunjungan/:id`
- **Response**:
  - **200 OK**: Returns kunjungan details.

##### Update Kunjungan

- **Endpoint**: `PUT /kunjungan/:id`
- **Request Body**: Same as Create Kunjungan
- **Response**:
  - **200 OK**: Kunjungan successfully updated.

##### Delete Kunjungan

- **Endpoint**: `DELETE /kunjungan/:id`
- **Response**:
  - **200 OK**: Kunjungan successfully deleted.
