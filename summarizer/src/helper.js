function serializerBook(book) {
  return book
    .replace(/<a[^>]*>.*<\/a>/g, "")
    .replace(/<img[^>]*>/g, "")
    .replace(/Recommended Reading.*$/is, "")
    .replace(/Other Books.*$/is, "")
    .replace(/<h2>Buy The Book.*$/is, "");
}

module.exports = {
  serializerBook,
};
