// console.log("Hello student");
// alert ("Good morning mother father");
//  document.write("Batch 47");

 //variable = 

 //VAR = bisa di deklarasikan ulang dan bisa diubah valuenya
 //LET = tidak bisa di deklarasikan ulang akan tetapi bisa diubah valuenya
 //CONST = tidak bisa di deklarasikan ulang dan tidak bisa diubah valuenya

 // tipe data ada 3. string integer dan boolean

//  let nama = "ibnu";
//  let umur = 24 ;
//  let isMarried = false ;

//  console.log('nama saya ${nama} umur saya ${umur} tahun');
//  console.log("nama saya", nama, "umur saya", umur, "tahun");
//  console.log("nama saya " + nama + " umur saya " + umur + " tahun ");

// pengkondisian

// let nilai = 80 ;

// if(nilai >= 75) {
//     console.log("kamu lulus!");
// } else {
//     console.log ("Kamu tidak lulus!");
// }

//function

// function Nama () {
//     let x = 5;
//     let y = 10;

//     let result = x * y;
//     console.log(result);
// }
// Nama();

// function Nama2 (x, y) {
//     console.log(x);
//     console.log(y);

//     let result = x * y;
//     console.log(result);
// }

// Nama2(5, 10);

// camelcase = namaSayaAdalah ( unwriten rules )

function submitData () {
    let name = document.getElementById("name").value;
    let email = document.getElementById("email").value;
    let phone = document.getElementById("number").value;
    let subject = document.getElementById("subject").value;
    let comment = document.getElementById("textarea").value;   
    
    if ( name == "" ) {
        return alert ("Nama tidak boleh dikosongkan") ;
    } else if ( email == "") {
        return alert ("Email tidak boleh dikosongkan");
    } else if ( phone == "" ) {
        return alert ("Nomor handphone tidak boleh dikosongkan");
    } else if ( subject == "") {
        return alert ("Subject tidak boleh dikosongkan") ;
    } else if ( comment == "") {
        return alert ("comment tidak boleh dikosongkan");
    }

    let emailReceiver = "ibnuhakim81@gmail.com";

    let a = document.createElement("a");
    a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya ${name}, ${comment}, Silahkan isi kontak saya di nomor ${phone}`;
    a.click();

    console.log(name);
    console.log(email);
    console.log(phone);
    console.log(subject);
    console.log(comment);

let emailer = {
    name,
    email,
    phone,
    subject,
    comment,
};

    console.log(emailer);

}