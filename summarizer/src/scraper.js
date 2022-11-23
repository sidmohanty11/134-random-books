const cheerio = require("cheerio");
const axios = require("axios");

const { serializerBook } = require("./helper");

async function scrapeBook({ title, idx }) {
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

module.exports = {
  scrapeBook,
};
