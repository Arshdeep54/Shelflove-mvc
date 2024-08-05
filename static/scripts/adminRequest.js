const adminCheckboxes = document.querySelectorAll(".admin-checkbox");
      const selectedUserIds = [];

      adminCheckboxes.forEach((checkbox) => {
        checkbox.addEventListener("change", (event) => {
          if (event.target.checked) {
            selectedUserIds.push(event.target.dataset.userId);
          } else {
            const index = selectedUserIds.indexOf(event.target.dataset.userId);
            if (index > -1) {
              selectedUserIds.splice(index, 1);
            }
          }
        });
      });
      const submitAdminButton =
        document.getElementsByClassName("approve-admin-btn");

      submitAdminButton[0].addEventListener("click", (event) => {
        event.preventDefault();
        if (selectedUserIds.length === 0) {
          alert("Please select at least one checkbox");
          return;
        }

        const requestBody = JSON.stringify({ ids: selectedUserIds });
        console.log(requestBody);

        fetch("/api/admin/approveadmin/", {
          method: "POST",
          body: requestBody,
          headers: { "Content-Type": "application/json" },
        }).then((response) => {
          selectedUserIds.length = 0;
          location.reload();
        });
      });