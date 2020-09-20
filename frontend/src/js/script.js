const baseUrl = 'http://localhost:8000'

function getAllArticles() {
  $.get(`${baseUrl}/articles`, data => {
    console.log(data)
  });
}

function jqReady() {
  $('.title').text("Hello from jQuery")

  getAllArticles()
}

$(jqReady)
