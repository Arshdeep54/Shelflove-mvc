const deleteBookButton = document.getElementById("deletebook");
const deletebookForm = document.getElementById("deletebookForm");
deleteBookButton.addEventListener("click", () => {
  deletebookForm.submit();
});
const today = new Date().toISOString().slice(0, 10);
document.getElementById("publication_date").value = today;
const updateBookButton = document.getElementById("updatebook");
const updatebookForm = document.getElementById("updatebookForm");

updateBookButton.addEventListener("click", () => {
  updatebookForm.submit();
});
