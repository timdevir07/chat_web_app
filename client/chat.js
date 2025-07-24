let username = localStorage.getItem("username");
if (!username) {
  alert("Please login first.");
  window.location.href = "login.html";
}

document.getElementById("welcome-text").textContent = `Logged in as: ${username}`;

const partnerInput = document.getElementById("partner-name");
const messageInput = document.getElementById("message-input");
const sendBtn = document.getElementById("send-btn");
const chatBox = document.getElementById("chat-box");

let currentPartner = "";

function sendMessage() {
  const content = messageInput.value.trim();
  const partner = partnerInput.value.trim();
  if (!content || !partner) return;

  fetch("http://localhost:8080/api/send", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ sender: username, receiver: partner, content })
  })
  .then(() => {
    messageInput.value = "";
    loadMessages(); // Refresh chat
  });
}

function loadMessages() {
  const partner = partnerInput.value.trim();
  if (!partner) return;

  if (partner !== currentPartner) {
    currentPartner = partner;
    chatBox.innerHTML = ""; // Clear on new partner
  }

  fetch(`http://localhost:8080/api/messages?user1=${username}&user2=${partner}`)
    .then(res => res.json())
    .then(data => {
      chatBox.innerHTML = "";
      data.forEach(msg => {
        const div = document.createElement("div");
        div.classList.add("message");
        div.classList.add(msg.sender === username ? "sender" : "receiver");
        div.textContent = `${msg.sender}: ${msg.content}`;
        chatBox.appendChild(div);
      });
      chatBox.scrollTop = chatBox.scrollHeight;
    });
}

// Event listeners
sendBtn.addEventListener("click", sendMessage);
messageInput.addEventListener("keypress", e => {
  if (e.key === "Enter") sendMessage();
});
partnerInput.addEventListener("change", loadMessages);
partnerInput.addEventListener("keyup", loadMessages);

setInterval(loadMessages, 2000);

// Logout
function logout() {
  localStorage.removeItem("username");
  window.location.href = "login.html";
}
