const cheerio = require("cheerio");
const axios = require("axios");

const { BOOKS } = require("./books");
const { serializerBook } = require("./helper");

async function getBook({ title, idx }) {
  let html = await axios.get(
    "https://www.samuelthomasdavies.com/book-summaries/"
  );
  let $ = cheerio.load(html.data);

  let link;

  $("ul").each((_, el) => {
    const heading = $(el).find("li > a").text();
    if (heading.includes(title)) {
      link = $(el).find("li > a").attr("href");
    }
  });

  html = await axios.get(link);
  $ = cheerio.load(html.data);

  let book = $("div.entry-content").html();
  book = serializerBook(book);

  $ = cheerio.load(book);
  const plainText = $(book).text();

  const data = {
    title,
    book,
    plainText,
    index: idx,
  };

  return data;
}

function getRandomBookFromArray() {
  const idx = Math.ceil(Math.random() * 135);
  return {
    title: BOOKS[idx],
    idx,
  };
}

async function fetchRandomBook() {
  return getBook(getRandomBookFromArray());
}

module.exports = {
  fetchRandomBook,
};
