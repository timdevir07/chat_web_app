document.getElementById("loginForm").addEventListener("submit", function (e) {
  e.preventDefault();

  const email = document.getElementById("email").value.trim();
  const password = document.getElementById("password").value.trim();

  if (!email || !password) {
    alert("Please fill all fields.");
    return;
  }

  fetch("http://localhost:8080/api/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password })
  })
    .then(res => res.json())
    .then(data => {
      if (data.message === "Login successful") {
        localStorage.setItem("username", data.name); // Get from backend response
        window.location.href = "chat.html";
      } else {
        alert(data.message || "Login failed");
      }
    })
    .catch(err => {
      alert("Something went wrong");
      console.error(err);
    });
});
