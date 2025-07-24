// client/script.js

const formTitle = document.getElementById('form-title');
const toggleText = document.getElementById('toggle-text');
const toggleLink = document.getElementById('toggle-link');
const submitBtn = document.getElementById('submit-btn');
let isLogin = true;

toggleLink.addEventListener('click', () => {
  isLogin = !isLogin;

  formTitle.textContent = isLogin ? 'Login' : 'Register';
  submitBtn.textContent = isLogin ? 'Login' : 'Register';
  toggleText.innerHTML = isLogin
    ? `Don't have an account? <span id="toggle-link">Register here</span>`
    : `Already have an account? <span id="toggle-link">Login here</span>`;

  // reattach event listener after innerHTML change
  document.getElementById('toggle-link').addEventListener('click', () => {
    toggleLink.click();
  });
});

// Later: Add fetch('/api/auth/login') or fetch('/api/auth/register') here
