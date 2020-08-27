# warung-makan-gerin

===== Cara Pakai backend yang ga jelas ini =======

1. Klick kanan dekstop -> git bash here -> ketik git clone https://github.com/gerins/warung-makan-gerin.git
2. Nanti didalem folder warung-makan-gerin nya, klick kanan -> git bash here -> ketik go run app.go
3. Jangan lupa sesuaikan file config.txt nya
4. Masuk ke folder Query & Routing -> Jalankan Warung_Makan_Gerin.sql nya biar kamu punya databasenya (kalo gatau tanya valdy)

5. Untuk post gambar, Route nya POST http://localhost:8080/menu 
6. Buka Postman -> Pilih Method POST abis itu paste link nya -> Pilih Body -> form-data -> 
7. isi key nya file (key nya wajib file biar bisa dibongkar di backend) -> value nya ya file nya itu
8. https://github.com/gerins/warung-makan-gerin/blob/master/domains/menu/menuController.go liat line 71, syntax untuk ambil file nya

Sisa route nya liat di file Routing.txt di Folder Query & Routing
*Kalo mau post menu harus post kategori menu dulu
