# agn-aggregator-toolbox
Aplikasi toolbox untuk aggregator. Aplikasi ini mengimplementasikan Hexagonal architecture agar lebih mudah dalam scaling, maintain, dan testing.

## Project Structure
.
├── cmd
├── db
│   ├── migration
│   └── query
├── internal
│   ├── adapter
│   │   ├── fiber
│   │   ├── mariadb
│   ├── application
│   │   └── domain
│   │       └── model
│   └── port
├── templates
│   └── fiber
│       ├── pages
│       │   ├── partner-produk
│       │   ├── product
│       │   ├── provider
│       │   └── transaksi-logs
│       └── partials
├── tmp
└── util

### Folder Legend
1. **/cmd** Merupakan folder yang berisi package main untuk bootstrap aplikasi, seperti setting adapter dan service (contoh, database adapter).
2. **/db** Merupakan folder yang berisi file yang berkaitan dengan sqlc.
    1. **/db/migration** Berisi file schema database.
    2. **/db/query** Berisi file-file untuk menggenarate file *.sql.go.
3. **/internal** Merupakan folder berisi core aplikasi. Terdapat 3 folder utama di dalam folder ini yaitu
    1. **adapter** Merupakan kumpulan adapter-adapter (implementasi dari port) seperti adapter untuk akses ke database atau http router (fiber).
    2. **application** Merupakan core logic aplikasi.
    3. **port** Meurpakan file-file interface(contract) yang akan menghubungkan core aplikasi dengan adapter
4. **/template** Berisi file-file html untuk keperluan rendering aplikasi. Termasuk file component untuk htmx.
5. **/util** Berisi fungsi-fungsi bantuan dan config.
