Loyihaning umumiy maqsadi:
Bu loyihaning maqsadi tenderlar (tovar yoki xizmatlar xarid qilish jarayoni) va ularga takliflarni boshqarish uchun backend tizimi qurishdir. Mijozlar tenderlarni e'lon qiladi, pudratchilar takliflar beradi, va tizim mijozlarga takliflarni ko'rib chiqish va shartnomalarni tanlangan pudratchilarga topshirish imkonini beradi. Loyihada foydalanuvchi autentifikatsiyasi, rol boshqaruvi, tender yaratish, takliflar yuborish va ularni baholash funksiyalari amalga oshiriladi.

Funksional imkoniyatlar:
1. Foydalanuvchi autentifikatsiyasi va rol boshqaruvi
Tavsifi:

Ro‘yxatdan o‘tish va tizimga kirish funksiyalari amalga oshiriladi.
Foydalanuvchilar mijoz yoki pudratchi bo‘lishi mumkin.
Mijozlar tender yaratadi, pudratchilar esa taklif beradi.
Autentifikatsiya JWT yoki sessiya tokenlari orqali amalga oshiriladi.
Rolga asoslangan boshqaruv: faqat mijozlar tender yaratishi, va faqat pudratchilar taklif yuborishi mumkin.
Talablar:

POST /register (Ro'yxatdan o'tish uchun endpoint).
POST /login (Tizimga kirish uchun endpoint).
Marshrutlarni himoya qilish uchun middleware (rolni tekshirish).
Baholash mezonlari:

Autentifikatsiyaning xavfsizligi.
Rolga asoslangan boshqaruv to‘g‘ri ishlashi.
2. Tender yaratish (faqat mijozlar tomonidan)
Tavsifi:

Mijozlar tender yaratishi kerak: sarlavha, tavsif, muddat va byudjet ko‘rsatiladi.
Tenderga loyiha xususiyatlari kiritilgan PDF kabi fayl qo‘shilishi mumkin.
Tenderning holati (masalan: "ochiq", "yopiq", "topshirilgan") bo‘lishi kerak.
Mijozlar o‘zlarining barcha tenderlari va ularning holatlarini ko‘rishi mumkin.
Talablar:

POST /tenders (Tender yaratish).
GET /tenders (Tenderlarni ro‘yxat qilish).
PUT /tenders/
(Tender holatini yangilash).
DELETE /tenders/
(Tenderni o‘chirish).
Baholash mezonlari:

Tender yaratish va boshqarish imkoniyati.
Tender yaratishda to‘g‘ri validatsiya (masalan, byudjet > 0, muddat kelajakda bo‘lishi kerak).
3. Taklif yuborish (faqat pudratchilar tomonidan)
Tavsifi:

Pudratchilar ochiq tenderlarga taklif yuborishi kerak.
Taklifda taklif qilingan narx, yetkazib berish muddati va ixtiyoriy izohlar bo‘ladi.
Pudratchilar barcha tenderlarni ko‘rishi va ularga taklif yuborishi mumkin.
Taklif yuborilganda mijozga xabar yuboriladi.
Talablar:

POST /tenders/
/bids (Taklif yuborish).
GET /tenders/
/bids (Yuborilgan takliflarni ko‘rish).
Baholash mezonlari:

Pudratchilar taklif yuborishi mumkinligi.
Taklif ma'lumotlarini validatsiya qilish (masalan, narx > 0, yetkazib berish muddati > 0).
4. Takliflarni filtrlash va saralash
Tavsifi:

Mijozlar takliflarni narx va yetkazib berish muddati kabi mezonlar bo‘yicha filtrlashi va saralashi mumkin.
Talablar:

GET /tenders/
/bids?price=<>&delivery_time=<> (Filtrlash imkoniyati).
Takliflarni narx yoki muddat bo‘yicha saralash.
5. Takliflarni baholash va tenderni topshirish
Tavsifi:

Tender muddati tugagandan so‘ng, mijoz barcha takliflarni baholashi va g‘olib taklifni tanlashi kerak.
G‘olib tanlanganda pudratchiga xabar beriladi.
Tenderning holati avtomatik ravishda "yopiq" yoki "topshirilgan" holatga o‘tadi.
Talablar:

POST /tenders/
/award/
(Tenderni topshirish).
Pudratchilarga xabar berish.
Baholash mezonlari:

Mijoz takliflarni ko‘rib chiqishi, tenderni topshirishi va pudratchiga xabar yuborishi.
Tenderning holati to‘g‘ri yangilanishi.
6. Tender va taklif tarixini ko‘rish
Tavsifi:

Mijozlar va pudratchilar o‘zlarining tenderlari va takliflarining tarixini ko‘rishi mumkin.
Talablar:

GET /users/
/tenders (Foydalanuvchining tenderlar tarixi).
GET /users/
/bids (Pudratchining takliflari tarixi).
Baholash mezonlari:

Har ikkala rol uchun tarixni ko‘rish imkoniyati.
Qo‘shimcha (bonus) funksiyalar:
Real-time yangilanishlar:

WebSocket orqali foydalanuvchilarga real vaqt rejimida xabar yuborish.
Rate limiting:

Pudratchilar bir daqiqada faqat 5 ta taklif yuborishi mumkin.
Keshlash:

Tenderlar va takliflar ro‘yxati kabi tez-tez ishlatiladigan ma'lumotlarni keshlash.
Ma'lumotlar bazasi sxemasi:
User: (id, username, password, role, email)
Tender: (id, client_id, title, description, deadline, budget, status)
Bid: (id, tender_id, contractor_id, price, delivery_time, comments, status)
Notification: (id, user_id, message, relation_id, type, created_at)
Texnik talablar:
Backend: Go (Golang).
Ma'lumotlar bazasi: PostgreSQL yoki MongoDB.
API: REST API.
Hujjatlar: Swagger.
Konteynerizatsiya: Docker.
Real vaqt: WebSocket.
Loyihani ishga tushirish:
Ma'lumotlar bazasi:
bash
Copy code
make run_db
Ilova:
bash
Copy code
make run


sqlc-generate:
  sqlc vet ; sqlc generate