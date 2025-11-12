# 🌱 AgriChain Transparency Platform

AgriChain Transparency Platform adalah aplikasi pertanian berbasis **Blockchain (Hyperledger Besu)** yang berfokus pada **transparansi alur distribusi hasil pertanian**.  
Platform ini tidak hanya mempertemukan petani, pengepul, pengolah, tengkulak, retailer, dan consumer, tetapi juga mencatat setiap proses distribusi secara **aman, transparan, dan tidak dapat dimanipulasi** melalui teknologi blockchain.
Terdapat juga fitur Markerplace untuk aktivitas jual beli online yang aman dan modern (ongoing).  

## ✨ Fitur Utama
- 📦 **Traceability**: Melacak alur distribusi hasil pertanian dari petani hingga konsumen akhir.  
- 🔗 **Integrasi Blockchain**: Menggunakan **Hyperledger** untuk transparansi data distribusi.  
- 👩‍🌾 **Multi-Role Access**: Mendukung role **Farmer,Collector, Processor, Tengkulak/Distributor, Retailer, dan Consumer**.  
- 📊 **Transparansi Harga & Distribusi**: Mengurangi asimetri informasi antara petani dengan tengkulak, dan tengkulak ke penjual.  
- ⚙️ **Clean Architecture**: Struktur kode modular, mudah dikembangkan, dan scalable.  
- 📜 **Dokumentasi Swagger**: Dokumentasi dengan Swagger untuk mempermudah testing dan Quality. Assurance.  


## 🏗️ Arsitektur
Aplikasi ini menggunakan pendekatan **Clean Architecture** dengan pemisahan layer:
- **Entity**: Database core.
- **Handler**: Interaksi pada request dan response, dan validasi awal.  
- **Use Case**: Logika aplikasi & interaksi third party.  
- **Repository**: Operasi database SQL & nonSQL. 
- **Infrastructure**: Database (MySQL), Blockchain (Hyperledger Besu).  

## 🛠️ Teknologi
- **Backend**: [Go (Golang)](https://go.dev/)  
- **Database**: [MySQL](https://www.mysql.com/)  
- **Blockchain**: [Hyperledger Besu](https://besu.hyperledger.org/)  
- **Architecture**: Clean Architecture.  
- **Redis**: [Redis](https://redis.io/)