const HOST = process.env.HOST || "http://localhost";
const PORT = process.env.PORT || 8050;

const express = require("express");
const { fetchRandomBook } = require("./scraper");

const app = express();

app.get("/books/random", async (req, res) => {
  const book = await fetchRandomBook();
  res.send({
    data: book,
  });
});

app.listen(8050, () => {
  console.log(`Listening at ${HOST}:${PORT}`);
});
