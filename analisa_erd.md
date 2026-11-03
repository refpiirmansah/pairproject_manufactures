## ERD Title: Order & Assembly Management System

1. Entities and Their Attributes:

## Attributes:

A. Entity: users
- Attributes:
  - id (PK, INT)
  - email (VARCHAR(255), UNIQUE)
  - password_hash (VARCHAR(255))

B. Entity: user_profiles
- Attributes:
  - id (PK, INT)
  - user_id (FK, INT, UNIQUE)
  - full_name (VARCHAR(100))
  - phone (VARCHAR(20))

C. Entity: customers
- Attributes:
  - id (PK, INT)
  - name (VARCHAR(100))
  - email (VARCHAR(255))
  - phone (VARCHAR(20))
  - address (TEXT)

D. Entity: orders
- Attributes:
  - id (PK, INT)
  - customer_id (FK, INT)
  - order_date (DATETIME)
  - status (VARCHAR(30))

E. Entity: order_items
- Attributes:
  - id (PK, INT)
  - order_id (FK, INT)
  - product_id (FK, INT)
  - quantity (INT)
  - unit_price (DECIMAL(12,2))

F. Entity: products
- Attributes:
  - id (PK, INT)
  - name (VARCHAR(100))
  - description (TEXT)
  - price (DECIMAL(12,2))
  - stock (INT)

G. Entity: components
- Attributes:
  - id (PK, INT)
  - name (VARCHAR(100))
  - category (VARCHAR(50))
  - stock (INT)
  - price (DECIMAL(12,2))

H. Entity: product_components
- Attributes:
  - id (PK, INT)
  - product_id (FK, INT)
  - component_id (FK, INT)
  - quantity (INT)

I. Entity: assemblies
- Attributes:
  - id (PK, INT)
  - order_item_id (FK, INT)
  - quantity (INT)
  - start_date (DATETIME)
  - end_date (DATETIME)
  - status (VARCHAR(30))

## Relationships:

- users to user_profiles
  A. Type: One to One (1:1)
  Deskripsi: Setiap pengguna (`users`) memiliki tepat satu profil (`user_profiles`), dan setiap profil terhubung secara unik ke satu pengguna.

- customers to orders
  B. Type: One to Many (1:N)
  Deskripsi: Satu pelanggan (`customers`) dapat memiliki banyak pesanan (`orders`), tetapi setiap pesanan hanya dimiliki oleh satu pelanggan.

- orders to order_items
  C. Type: One to Many (1:N)
  Deskripsi: Satu pesanan (`orders`) dapat berisi banyak item pesanan (`order_items`), tetapi setiap item pesanan hanya terikat pada satu pesanan.

- products to order_items
  D. Type: One to Many (1:N)
  Deskripsi: Satu produk (`products`) dapat dipesan dalam banyak item pesanan (`order_items`), tetapi setiap item pesanan mencatat satu produk spesifik.

- products to product_components
  E. Type: One to Many (1:N)
  Deskripsi: Satu produk (`products`) dapat menggunakan banyak komponen melalui entitas perantara `product_components`.

- components to product_components
  F. Type: One to Many (1:N)
  Deskripsi: Satu komponen (`components`) dapat digunakan dalam pembuatan banyak produk melalui entitas perantara `product_components`.

- order_items to assemblies
  G. Type: One to Many (1:N)
  Deskripsi: Satu item pesanan (`order_items`) dapat memiliki satu atau lebih jadwal/proses perakitan (`assemblies`).

## Integrity Constraints:

- Email pada tabel `users` harus bersifat unik (UNIQUE).
- Foreign key `user_id` pada tabel `user_profiles` bersifat UNIQUE untuk menjaga tingkat keakuratan relasi 1:1.
- Nilai `price` pada tabel `products` & `components`, serta `unit_price` pada `order_items` harus berupa nilai numerik desimal non-negatif (`DECIMAL(12,2)`).
- Nilai `quantity` dan `stock` harus berupa angka bulat positif (non-negatif).
- Kolom `status` pada tabel `orders` terbatas pada opsi status berikut: `pending`, `waiting_stock`, `assembling`, `ready`, `completed`, `rejected`.

## Additional Notes:

- Tabel `product_components` bertindak sebagai *junction table* (penghubung) untuk memecah relasi Many to Many (N:M) antara `products` dan `components`.
- Tabel `order_items` bertindak sebagai penampung detail produk yang dibeli beserta harga per unit pada saat transaksi dilakukan.
- Tabel `assemblies` mencatat progres perakitan fisik barang berdasarkan item yang telah dipesan oleh pelanggan.