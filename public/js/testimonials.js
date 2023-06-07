// // step:
// // - making class
// // - making object
// // - called it so it appears in browser using loop and innerHTML
// // - refactor the code so it implements inheritance AuthorTestimonial
// // - refactor the code so it implements polymorphism by making class CompanyTestimonial and override get author(the key point is polymorhp implement and override)
// // - refactor the code so it implements abstraction by making get testimonialHTML and get author throw error when the children didn't use it
// // - refactor the code so it implements encapsulation, use setter and getter (key point is, encapsulation try to make app more secure by implementing restriction)

// // Superclass
// class Testimonial {
//     #quote = "";
//     #image = "";
  
//     constructor(quote, image) {
//       this.#quote = quote;
//       this.#image = image;
//     }
  
//     get quote() {
//       return this.#quote;
//     }
  
//     get image() {
//       return this.#image;
//     }
  
//     // This is an abstract method that subclasses will implement
//     get author() {
//       throw new Error("getAuthor() method must be implemented");
//     }
  
//     // This is a polymorphic method that can take any subclasses of Testimonial
//     get testimonialHTML() {
//       return `<div class="testimonial">
//                   <img
//                       src="${this.image}"
//                       class="profile-testimonial"
//                   />
//                   <p class="quote">${this.quote}</p>
//                   <p class="author">- ${this.author}</p>
//               </div>
//           `;
//     }
//   }
  
//   // Subclass
//   class AuthorTestimonials extends Testimonial {
//     #author = "";
  
//     constructor(author, quote, image) {
//       super(quote, image);
//       this.#author = author;
//     }
  
//     get author() {
//       return this.#author;
//     }
//   }
  
//   // Subclass
//   class CompanyTestimonials extends Testimonial {
//     #company = "";
  
//     constructor(company, quote, image) {
//       super(quote, image);
//       this.#company = company;
//     }
  
//     get author() {
//       return this.#company + " Company";
//     }
//   }
  
//   const testimonial1 = new AuthorTestimonials(
//     "Ibnu hakim",
//     "DGP Riders",
//     "https://static.wikia.nocookie.net/8fe0e2eb-ffab-4981-bd91-6b77a93f33dc/scale-to-width/755"
//     );
//   const testimonial2 = new AuthorTestimonials(
//     "Winter",
//     "AESPA Hyper LINKS",
//     "https://nuunuu.org/wp-content/uploads/2023/01/Winter-aespa-featured.jpeg"
//     );
//   const testimonial3 = new CompanyTestimonials(
//     "BMTH",
//     "Let's Rock!!! 🔥🔥🔥",
//     "https://cdn-2.tstatic.net/tribunnews/foto/bank/images/Bring-Me-The-Horizon.jpg"
//     );
  
//   let testimonialData = [testimonial1, testimonial2, testimonial3];
//   let testimonialHTML = "";
  
//   for (let i = 0; i < testimonialData.length; i++) {
//     testimonialHTML += testimonialData[i].testimonialHTML;
//   }
  
//   document.getElementById("testimonials").innerHTML = testimonialHTML;

// FILTERIN

// const testimonialData = [
//   {
//     author: "Ibnu hakim",
//     quote: "Jagalah Kebersihan",
//     image: "https://static.wikia.nocookie.net/8fe0e2eb-ffab-4981-bd91-6b77a93f33dc/scale-to-width/755",
//     rating: 5,
//   },
//   {
//     author: "Winter",
//     quote: "AESPA Hyper LINKS",
//     image:"https://nuunuu.org/wp-content/uploads/2023/01/Winter-aespa-featured.jpeg",
//     rating: 4,
//   },
//   {
//     author: "BMTH",
//     quote: "Lets ROCK !!!",
//     image: "https://cdn-2.tstatic.net/tribunnews/foto/bank/images/Bring-Me-The-Horizon.jpg",
//     rating: 4,
//   },

//   {
//     author: "Alucard",
//     quote: "hey hey not bad",
//     image: "https://media.hitekno.com/thumbs/2019/12/03/86483-hero-alucard-dari-mobile-legends/730x480-img-86483-hero-alucard-dari-mobile-legends.jpg",
//     rating: 3,
//   },

//   {
//     author: "Kono kobo",
//     quote: "Tydack ramah bintang 1",
//     image: "https://pic-bstarstatic.akamaized.net/ugc/0e27f3abe4cb92a87a7d28fc723b19fe.jpg@1200w_630h_1e_1c_1f.webp",
//     rating: 1,
//   },
// ];

// function allTestimonials() {
//   let testimonialHTML = "";

//   testimonialData.forEach(function (item) {
//     testimonialHTML += `<div class="testimonial">
//                             <img
//                                 src="${item.image}"
//                                 class="profile-testimonial"
//                             />
//                             <p class="quote">${item.quote}</p>
//                             <p class="author">- ${item.author}</p>
//                             <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
//                         </div>
//         `;
//   });

//   document.getElementById("testimonials").innerHTML = testimonialHTML;
// }

// allTestimonials();

// function filterTestimonials(rating) {
//   let testimonialHTML = "";

//   const testimonialFiltered = testimonialData.filter(function (item) {
//     return item.rating === rating;
//   });

//   //   console.log(testimonialFiltered);

//   if (testimonialFiltered.length === 0) {
//     testimonialHTML += `<h1>Data not found!</h1>`;
//   } else {
//     testimonialFiltered.forEach(function (item) {
//       testimonialHTML += `<div class="testimonial">
//                                 <img
//                                     src="${item.image}"
//                                     class="profile-testimonial"
//                                 />
//                                 <p class="quote">${item.quote}</p>
//                                 <p class="author">- ${item.author}</p>
//                                 <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
//                             </div>
//                         `;
//     });
//   }

//   document.getElementById("testimonials").innerHTML = testimonialHTML;
// }