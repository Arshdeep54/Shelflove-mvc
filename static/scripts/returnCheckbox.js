const returnCheckboxes = document.querySelectorAll(".return-checkbox");
const selectedIssueIds = [];
let selectedReturnbooks = {};

returnCheckboxes.forEach((checkbox) => {
  checkbox.addEventListener("change", (event) => {
    if (event.target.checked) {
      selectedIssueIds.push(event.target.dataset.issueId);

      if (!selectedReturnbooks[event.target.dataset.bookId]) {
        selectedReturnbooks[event.target.dataset.bookId] = 1;
      } else {
        selectedReturnbooks[event.target.dataset.bookId]++;
      }
    } else {
      const index = selectedIssueIds.indexOf(event.target.dataset.issueId);
      if (index > -1) {
        selectedIssueIds.splice(index, 1);
      }
      selectedReturnbooks[event.target.dataset.bookId]--;
    }
  });
});

const submitButton = document.getElementsByClassName("approve-btn");
submitButton[0].addEventListener("click", (event) => {
  event.preventDefault();
  if (selectedIssueIds.length === 0) {
    alert("Please select at least one checkbox");
    return;
  }

  const requestBody = JSON.stringify({
    issueIds: selectedIssueIds,
    selectedBooks: selectedReturnbooks,
  });
  console.log(requestBody);
  fetch("/api/admin/approvereturns", {
    method: "POST",
    body: requestBody,
    headers: { "Content-Type": "application/json" },
  }).then((response) => {
    selectedIssueIds.length = 0;
    selectedReturnbooks = {};
    location.reload();
  });
});
