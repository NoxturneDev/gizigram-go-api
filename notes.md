# first commit

# Project Structure

entry file: main.go

## Directory and packages
main package:
  routes.go - contains all the routes
  util.go - for all the utility functions (like logging)

handlers/ - contains all the hadnlers
  ../users.go

model/ - contains all the models struct
  ../user.go

services/
  ../user.go - contains all the business logic 

-- gender
1 -> laki-laki
0 -> perempuan

localhost:3002/api/users/
--
{
"username": "bapak",
"password": "anak"
}

localhost:3002/api/users/
--
{
"username": "bapak",
"password": "anak"
}



localhost:3002/api/parent/create
--
{
"name" : "ibu 2",
"user_id": 1
}

localhost:3002/api/children/create
--
{
"name": "jerry",
"age" : 20,
"height": "178 cm",
"parent_id": 2
}

localhost:3002/api/login
--
{
"username" : "ibu",
"password" : "anak"
}


localhost:3002/api/parent/create
--
{
"name" : "ibu 2",
"user_id": 1
}

localhost:3002/api/children/create
--
{
"name": "jerry",
"age" : 20,
"height": "178 cm",
"parent_id": 2
}

localhost:3002/api/login
--
{
"username" : "ibu",
"password" : "anak"
}


prompt
-- 
-- Chat\
Rancang aplikasi yang ramah pengguna untuk menganalisis data anak-anak dalam jangka waktu tertentu guna mengidentifikasi potensi stunting berdasarkan standar WHO, dengan fitur seperti entri data, analisis data, pelaporan, tindakan pencegahan, manajemen pengguna, aksesibilitas mobile, dan dukungan serta umpan balik, serta integrasikan data WHO tentang nutrisi, perawatan kesehatan, kebersihan, sanitasi, pendidikan, dan kesadaran, untuk memberikan rekomendasi pencegahan stunting dan isyarat visual ketika metrik pertumbuhan anak di bawah ambang batas WHO

-- Image\
Analisis gambar menggunakan teknologi pengenalan gambar AI untuk menjelaskan fakta nutrisi dari sebuah makanan secara lengkap, termasuk evaluasi kesesuaian nutrisi makanan tersebut berdasarkan data sentiment dari WHO tentang stunting untuk menentukan apakah makanan ini cocok bagi pengguna
62 176 210 130 157 30 242 155 231 31 51 228 117 246 183 25 31 255 217