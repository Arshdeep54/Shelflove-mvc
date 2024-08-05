const numberInput = document.getElementById("positiveQuantityInput");
const errorQuant = document.getElementById("errorQuant");

numberInput.addEventListener("keypress", function (event) {
  if (event.key === "Backspace" || event.key === "Delete") {
    errorQuant.innerHTML = " ";
    return;
  }

  const isNumber = !isNaN(event.key) && event.key !== ".";

  if (!isNumber) {
    event.preventDefault();
    errorQuant.innerHTML = `Can't add negative or decimal number`;
  } else if (
    event.key === "Backspace" ||
    event.key === "Delete" ||
    isNumber
  ) {
    errorQuant.innerHTML = " ";
  }
});