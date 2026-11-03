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
  - email (VARCHAR(255), UNIQUE)
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
  - quantity (INT, NOT NULL, CHECK: $quantity > 0$)
  - unit_price (DECIMAL(12,2), NOT NULL, CHECK: $unit\_price \ge 0$)

F. Entity: products

- Attributes:
  - id (PK, INT)
  - name (VARCHAR(100))
  - description (TEXT)
  - price (DECIMAL(12,2), NOT NULL, CHECK: $price \ge 0$)
  - stock (INT, NOT NULL, CHECK: $stock \ge 0$)

G. Entity: components

- Attributes:
  - id (PK, INT)
  - name (VARCHAR(100))
  - category (VARCHAR(50))
  - stock (INT, NOT NULL, CHECK: $stock \ge 0$)
  - price (DECIMAL(12,2), NOT NULL, CHECK: $price \ge 0$)

H. Entity: product_components

- Attributes:
  - id (PK, INT)
  - product_id (FK, INT)
  - component_id (FK, INT)
  - quantity (INT, NOT NULL, CHECK: $quantity > 0$)
  - Constraint: UNIQUE (product_id, component_id)

I. Entity: assemblies

- Attributes:
  - id (PK, INT)
  - order_item_id (FK, INT)
  - quantity (INT, NOT NULL, CHECK: $quantity > 0$)
  - start_date (DATETIME, NULL)
  - end_date (DATETIME, NULL)
  - status (VARCHAR(30))
  - Constraint: CHECK ($end\_date\text{ IS NULL OR }end\_date \ge start\_date$)

J. Entity: assembly_components

- Attributes:
  - id (PK, INT)
  - assembly_id (FK, INT)
  - component_id (FK, INT)
  - quantity_used (INT, NOT NULL, CHECK: $quantity\_used > 0$)
  - Constraint: UNIQUE (assembly_id, component_id)

## Relationships:

- users to user_profiles
  A. Type: One to One (1:1)
  Deskripsi: Setiap pengguna (`users`) memiliki tepat satu profil (`user_profiles`), dan setiap profil terhubung secara unik ke satu pengguna.

- customers to orders
  B. Type: One to Many (1:N)
  Deskripsi: Satu pelanggan (`customers`) dapat memiliki banyak pesanan (`orders`), tetapi setiap pesanan hanya dimiliki oleh satu pelanggan.

-orders to order_items
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

- assemblies to assembly_components
  H. Type: One to Many (1:N)
  Deskripsi: Satu proses perakitan (`assemblies`) dapat mencatat pemakaian banyak komponen melalui entitas perantara `assembly_components`.

- components to assembly_components
  I. Type: One to Many (1:N)
  Deskripsi: Satu jenis komponen (`components`) dapat dicatat penggunaan dan jumlahnya dalam berbagai proses perakitan melalui `assembly_components`.

## Integrity Constraints:

- Email pada tabel `users` dan `customers` harus bersifat unik (UNIQUE).
- Foreign key `user_id` pada tabel `user_profiles` bersifat UNIQUE untuk menjaga tingkat keakuratan relasi 1:1.
- Pasangan (`product_id`, `component_id`) pada `product_components` bersifat UNIQUE untuk mencegah duplikasi pemetaan komponen pada produk yang sama.
- Pasangan (`assembly_id`, `component_id`) pada `assembly_components` bersifat UNIQUE untuk mencegah pemetaan duplikat komponen pada proses perakitan yang sama.
- Nilai `price` dan `stock` pada tabel `products` serta `components` bernilai NOT NULL dan tidak boleh negatif ($price \ge 0$, $stock \ge 0$).
- Nilai `unit_price` pada `order_items` bernilai NOT NULL dan tidak boleh negatif ($unit\_price \ge 0$).
- Nilai `quantity` pada `order_items`, `product_components`, dan `assemblies` bernilai NOT NULL dan harus bernilai bulat positif ($quantity > 0$).
- Nilai `quantity_used` pada `assembly_components` bernilai NOT NULL dan harus bernilai lebih dari nol ($quantity\_used > 0$).
- Tanggal selesai perakitan (`end_date`) pada tabel `assemblies` harus bernilai NULL atau lebih besar/sama dengan tanggal mulai (`end_date \ge start_date`).
- Kolom `status` pada tabel `orders` terbatas pada opsi status berikut: `pending`, `waiting_stock`, `assembling`, `ready`, `completed`, `rejected`.

## Business Rules:

-**Pengurangan Stok Produk:** Nilai `products.stock` berkurang saat barang dikeluarkan dari gudang untuk diproses dalam pesanan.

- **Penambahan Stok Produk:** Nilai `products.stock` bertambah ketika hasil perakitan selesai dan lolos proses Quality Control (QC).

## Additional Notes:

- Tabel `assembly_components` bertindak sebagai _junction table_ yang mencatat rincian komponen aktual beserta kuantitas yang digunakan (`quantity_used`) dalam setiap sesi perakitan.
- Tabel `product_components` bertindak sebagai _junction table_ (Bill of Materials / BOM) untuk memecah relasi Many to Many ($N:M$) antara `products` dan `components`.
- Tabel `order_items` bertindak sebagai penampung detail produk yang dibeli beserta harga per unit pada saat transaksi dilakukan.
- Tabel `assemblies` mencatat progres perakitan fisik barang berdasarkan item yang telah dipesan oleh pelanggan.
