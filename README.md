# Olah Product Image

Program untuk mengubah file-file images di suatu direktori ke format product image transfashion web, sehingga bisa diakses dengan format

    https://product.transfashionindonesia.com/data/images/eag/150098/0060/1.jpg


## Format Gambar
Format dasar image dalam direktori diformat sebagai berikut

### Dengan kode color

    [article]~[col]~[sequence].jpg

Hasil eksekusi program akan meyimpan program dengan struktur

    [article]                   <dir>
        [col]                   <dir>
            [sequence].jpg      <file>
            1.jpg               <file>
            2.jpg               <file>
            3.jpg               <file>
            dst


### Tanpa kode color

    [article]~[sequence].jpg

Hasil eksekusi akan sebagai berikut

    article                     <dir>
        [sequence].jpg          <file>
        1.jpg                   <file>
        2.jpg                   <file>
        3.jpg                   <file>
        dst    

## Eksekusi Program

Dari root direktori, eksekusi dengan

    go run ./main --source="/direktori/sumber" --target="/direktori/tujuan"

Atau apabila menggunakan yang versi compiled

    ./productimage --source="/direktori/sumber" --target="/direktori/tujuan"


