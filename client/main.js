
async function fetchJson(url) {
  const response = await fetch(url)
  const jsonData = await response.json()
  return jsonData 
}

const books = document.querySelector('.books')

function renderBook(book) {
  const bookContainer = document.createElement("div")
  bookContainer.classList.add("bookContainer")

  const heading = document.createElement("div")
  //todo: sanitize strings
  heading.innerHTML = `<b>${book.originalTitle}</b>: ${book.authors}`

  const image = document.createElement("img");
  image.src = book.imageUrl;
  image.style.width = "100px";
  
  const ratingContainer = document.createElement("div")
  ratingContainer.classList.add("ratingsContainer")

  const rating = document.createElement("div")
  rating.textContent = parseFloat(book.averageRating).toFixed(1)
  rating.style.fontSize = "20px"
  rating.style.fontWeight = "700"
  ratingContainer.appendChild(rating)

  const averageRatingFloat = parseFloat(book.averageRating)
  const starsContainer = document.createElement("div")
  for (let i = 0; i < averageRatingFloat; i++) {
    const star = document.createElement("img")
    star.style.height = '17px'
    if (i > averageRatingFloat - 1 && averageRatingFloat % 1 !== 0) {
      star.src = './half_star.png'
    } else {
      star.src = './full_star.png'
    }
    starsContainer.appendChild(star)
  }
  ratingContainer.appendChild(starsContainer)

  bookContainer.appendChild(heading)
  bookContainer.appendChild(image)
  bookContainer.appendChild(ratingContainer)

  books.appendChild(bookContainer);
}

fetchJson("http://localhost:8080/books/").then(response => {
  if (response.length) {
    response.forEach(book => {
      renderBook(book)
    })
  }
}).catch(error => {
  console.log("Error: ", error)
})

const openUploadFormBtn = document.getElementById("openUploadFormBtn")
openUploadFormBtn.addEventListener('click', () => {
  document.getElementById("uploadFormDialog").showModal()
})

const bookProperties = [
  "authors",
  // "id",
  // "created_at",
  // "updated_at",
  // "deleted_at",
  // "book_id",
  // "isbn",
  // "originalPublicationYear",
  "originalTitle",
  "imageUrl",
  "averageRating",
  // "ratingCount",
  // "ratings1",
  // "ratings2",
  // "ratings3",
  // "ratings4",
  // "ratings5",
]