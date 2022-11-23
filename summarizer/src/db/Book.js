const mongoose = require("mongoose");

const bookSchema = new mongoose.Schema({
  title: {
    type: mongoose.Schema.Types.String,
  },
  book: {
    type: mongoose.Schema.Types.String,
  },
  plainText: {
    type: mongoose.Schema.Types.String,
  },
  index: {
    type: mongoose.Schema.Types.Number,
  },
});

const Book = mongoose.model("Book", bookSchema);

module.exports = {
  Book,
};
