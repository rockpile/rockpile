(ql:quickload "clsql-sqlite3")
(clsql:connect `("data.db" :database-type :sqlite3)
(clsql:disconnect)
