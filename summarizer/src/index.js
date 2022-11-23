process.env.NODE_ENV !== "production" && require("dotenv").config();

const HOST = process.env.HOST || "http://localhost";
const PORT = process.env.PORT || 8050;

const express = require("express");
const mongoose = require("mongoose");
const { fetchRandomBook } = require("./store");

const app = express();

mongoose
  .connect(process.env.MONGO_URI)
  .then(() => {
    console.log("Connected to the DB!!");
  })
  .catch((err) => {
    console.log(err);
  });

app.get("/books/random", async (req, res) => {
  const book = await fetchRandomBook();
  res.send({
    data: book,
  });
});

app.listen(8050, () => {
  console.log(`Listening at ${HOST}:${PORT}`);
});
