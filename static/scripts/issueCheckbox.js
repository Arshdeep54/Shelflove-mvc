const issueCheckboxes = document.querySelectorAll(".issue-checkbox");
const selectedIssueReqIds = [];
let selectedIssuebooks = {};

issueCheckboxes.forEach((checkbox) => {
  checkbox.addEventListener("change", (event) => {
    if (event.target.checked) {
      console.log("checked");
      let bookQuantity = event.target.dataset.bookQuantity;
      if (
        !selectedIssuebooks[event.target.dataset.bookId] &&
        bookQuantity > 0
      ) {
        console.log("can issue");
        selectedIssueReqIds.push(event.target.dataset.issueId);
        selectedIssuebooks[event.target.dataset.bookId] = 1;
        console.log(selectedIssueReqIds, selectedIssuebooks);
      } else {
        if (
          bookQuantity > selectedIssuebooks[event.target.dataset.bookId]
        ) {
          selectedIssueReqIds.push(event.target.dataset.issueId);
          selectedIssuebooks[event.target.dataset.bookId]++;
        } else {
          alert("Not enough books");
          event.target.checked = false;
        }
      }
    } else {
      const index = selectedIssueReqIds.indexOf(
        event.target.dataset.issueId
      );
      if (index > -1) {
        selectedIssueReqIds.splice(index, 1);
      }
      selectedIssuebooks[event.target.dataset.bookId]--;
    }
  });
});
const submitIssueButton =
  document.getElementsByClassName("approve-issue-btn");
submitIssueButton[0].addEventListener("click", (event) => {
  event.preventDefault();
  if (selectedIssueReqIds.length === 0) {
    alert("Please select at least one checkbox");
    return;
  }
  const requestBody = JSON.stringify({
    issueIds: selectedIssueReqIds,
    selectedBooks: selectedIssuebooks,
  });

  console.log(requestBody);
  fetch("/api/admin/approveissues/", {
    method: "POST",
    body: requestBody,
    headers: { "Content-Type": "application/json" },
  }).then((response) => {
    selectedIssueReqIds.length = 0;
    selectedIssuebooks = {};
    location.reload();
  });
});