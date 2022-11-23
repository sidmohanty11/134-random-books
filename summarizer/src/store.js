const { Book } = require("./db/Book");
const { BOOKS } = require("./books");
const { scrapeBook } = require("./scraper");

async function getBook({ index, title }) {
  const found = await Book.findOne({
    index,
  });

  if (!found) {
    const data = await scrapeBook({ index, title });
    await Book.create({
      title: data.title,
      book: data.book,
      plainText: data.plainText,
      index: index,
    });
    return data;
  }

  return found;
}

function getRandomBookFromArray() {
  const idx = Math.ceil(Math.random() * 135);
  return {
    title: BOOKS[idx],
    index: idx,
  };
}

async function fetchRandomBook() {
  return getBook(getRandomBookFromArray());
}

module.exports = {
  fetchRandomBook,
};
