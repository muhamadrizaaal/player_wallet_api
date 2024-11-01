# Player Wallet API

API untuk mengelola registrasi pemain, login, pengelolaan bank, dan dompet pada sistem `player_wallet_api`.

## Menjalankan Server

Pastikan database PostgreSQL dan Redis sudah berjalan. Lalu, gunakan perintah berikut untuk menjalankan server:

```bash
  go run cmd/main.go
```

## Daftar Endpoint

### 1. **Auth Endpoints**
   - **Register**: Membuat akun baru untuk pemain.
     - **URL**: `POST /api/v1/players/register`
     - **Request Body**:
       ```json
       {
         "username": "string",
         "password": "string",
       }
       ```
   - **Login**: Masuk ke akun pemain.
     - **URL**: `POST /api/v1/players/login`
     - **Request Body**:
       ```json
       {
         "username": "string",
         "password": "string"
       }
       ```

### 2. **Player Endpoints**
   - **Logout**: Mengeluarkan pemain dari akun.
     - **URL**: `POST /api/v1/players/logout`
     - **Header**: 
       - `Authorization: Bearer <token>`
   
   - **Get All Players**: Mendapatkan daftar pemain dengan berbagai filter.
     - **URL**: `GET /api/v1/players`
     - **Header**:
       - `Authorization: Bearer <token>`
     - **Query Parameters (Filters)**:
       - `username` (string): Menyaring berdasarkan username.
       - `bank_name` (string): Menyaring berdasarkan nama bank yang terhubung.
       - `account_name` (string): Menyaring berdasarkan nama pemilik rekening.
       - `account_number` (string): Menyaring berdasarkan nomor rekening.
       - `min_balance` (float): Menyaring berdasarkan saldo minimal.
       - `register_at` (date): Menyaring berdasarkan tanggal registrasi.

     **Contoh Request**:
     ```http
     GET /api/v1/players?username=johndoe&bank_name=BCA&min_balance=1000
     ```

   - **Get Player by ID**: Mendapatkan data pemain berdasarkan ID.
     - **URL**: `GET /api/v1/players/:id`
     - **Parameter Path**:
       - `id`: ID pemain
     - **Header**:
       - `Authorization: Bearer <token>`

### 3. **Bank Endpoints**
   - **Register Bank**: Menambahkan bank yang terhubung dengan akun pemain.
     - **URL**: `POST /api/v1/players/banks`
     - **Header**:
       - `Authorization: Bearer <token>`
     - **Request Body**:
       ```json
       {
         "bank_name": "string",
         "account_name": "string",
         "account_number": "string"
       }
       ```
   
   - **Get Player Banks**: Mendapatkan daftar bank yang terhubung ke akun pemain.
     - **URL**: `GET /api/v1/players/banks`
     - **Header**:
       - `Authorization: Bearer <token>`

### 4. **Wallet Endpoints**
   - **Top Up Wallet**: Melakukan top-up ke dompet pemain.
     - **URL**: `POST /api/v1/players/wallet/topup`
     - **Header**:
       - `Authorization: Bearer <token>`
     - **Request Body**:
       ```json
       {
         "amount": "float"
       }
       ```
   
   - **Get Wallet**: Mendapatkan informasi saldo dompet pemain.
     - **URL**: `GET /api/v1/players/wallet`
     - **Header**:
       - `Authorization: Bearer <token>`

## Detail Filter pada Endpoint Get All Players

Endpoint `GET /api/v1/players` memiliki beberapa filter untuk menyaring data pemain sesuai dengan parameter berikut:
- `username`: Menyaring daftar pemain yang memiliki username yang sesuai.
- `bank_name`: Menyaring daftar pemain berdasarkan nama bank yang terhubung dengan akun.
- `account_name`: Menyaring berdasarkan nama pemilik rekening yang terhubung.
- `account_number`: Menyaring berdasarkan nomor rekening yang terhubung.
- `min_balance`: Menyaring daftar pemain yang memiliki saldo minimum sesuai dengan nilai yang ditentukan.
- `register_at`: Menyaring berdasarkan tanggal pendaftaran akun pemain.

**Contoh Request dengan Beberapa Filter**:
```http
GET /api/v1/players?username=johndoe&bank_name=BCA&account_number=123456789
