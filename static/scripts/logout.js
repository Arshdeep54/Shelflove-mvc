const logoutButton = document.getElementById("logout");
logoutButton.addEventListener("click", async () => {
  try {
    const response = await fetch("/api/auth/logout", {
      method: "GET",
    });
    const data = await response.json();
    if (data.message === "Successfully logged out") {
      window.location.href = "/login";
    } else {
      error(data.message);
    }
  } catch (error) {
    alert(error);
  }
});