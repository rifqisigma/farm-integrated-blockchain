# 🌱 AgriChain Transparency Platform

AgriChain Transparency Platform adalah aplikasi pertanian berbasis **Blockchain (Hyperledger)** yang berfokus pada **transparansi alur distribusi hasil pertanian**.  
Platform ini tidak hanya mempertemukan petani, tengkulak, distributor, retailer, dan consumer, tetapi juga mencatat setiap proses distribusi secara **aman, transparan, dan tidak dapat dimanipulasi** melalui teknologi blockchain.
Terdapat juga fitur Markerplace untuk aktivitas jual beli online yang aman dan modern.  

## ✨ Fitur Utama
- 📦 **Traceability**: Melacak alur distribusi hasil pertanian dari petani hingga konsumen akhir.  
- 🔗 **Integrasi Blockchain**: Menggunakan **Hyperledger** untuk transparansi data distribusi.  
- 👩‍🌾 **Multi-Role Access**: Mendukung role **Farmer, Tengkulak/Distributor, Retailer, dan Consumer**.  
- 📊 **Transparansi Harga & Distribusi**: Mengurangi asimetri informasi antara petani dengan tengkulak, dan tengkulak ke penjual .  
- ⚙️ **Clean Architecture**: Struktur kode modular, mudah dikembangkan, dan scalable.  
- 📜 **Dokumentasi Swagger**: Dokumentasi dengan Swagger untuk mempermudah testing dan Quality Assurance.  

## 🏗️ Arsitektur
Aplikasi ini menggunakan pendekatan **Clean Architecture** dengan pemisahan layer:
- **Entity**: Database core.
- **Handler**: Interaksi pada request dan response, dan validasi awal.  
- **Use Case**: Logika aplikasi & interaksi domain.  
- **Repository**: Operasi database SQL & nonSQL. 
- **Infrastructure**: Database (MySQL), Blockchain (Hyperledger Fabric).  

## 🛠️ Teknologi
- **Backend**: [Go (Golang)](https://go.dev/)  
- **Database**: [MySQL](https://www.mysql.com/)  
- **Blockchain**: [Hyperledger Fabric](https://www.hyperledger.org/use/fabric)  
- **Architecture**: Clean Architecture  