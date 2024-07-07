class AuthFormComponent extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.isLogin = true;
  }

  connectedCallback() {
    this.render();
    this.addEventListeners();
  }

  render() {
    const formType = this.isLogin ? 'Login' : 'Register';
    const switchText = this.isLogin ? 'Register here' : 'Login here';
    const switchId = this.isLogin ? 'switch-to-register' : 'switch-to-login';

    this.shadowRoot.innerHTML = `
      <gowc-card>
        <gowc-card-header title="${formType}"></gowc-card-header>
        <form id="auth-form">
          <gowc-input type="text" name="username" placeholder="Username"></gowc-input>
          ${!this.isLogin ? '<gowc-input type="email" name="email" placeholder="Email"></gowc-input>' : ''}
          <gowc-input type="password" name="password" placeholder="Password"></gowc-input>
          <button type="submit">${formType}</button> <!-- TODO: Fix button component -->
        </form>
        <p>${this.isLogin ? 'No account?' : 'Already have an account?'} <a href="#" id="${switchId}">${switchText}</a></p>
      </gowc-card>
    `;
  }

  addEventListeners() {
    const form = this.shadowRoot.querySelector('#auth-form');
    const switchLink = this.shadowRoot.querySelector('#switch-to-register, #switch-to-login');

    form.addEventListener('submit', this.handleSubmit.bind(this));
    switchLink.addEventListener('click', this.toggleForm.bind(this));
  }

  async handleSubmit(event) {
    event.preventDefault();
    const formData = new FormData(event.target);
    const endpoint = this.isLogin ? '/login' : '/register';

    try {
      const response = await fetch(endpoint, { /** TODO: Abstract into Controller */
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams(formData),
      });

      const result = await response.json();
      if (response.ok) {
        alert(this.isLogin ? 'Login successful' : 'Registration successful');
        if (this.isLogin) {
          localStorage.setItem('token', result.token);
        }
      } else {
        alert(result.message || 'An error occurred');
      }
    } catch (error) {
      console.error('Error:', error);
      alert('An error occurred. Please try again.');
    }
  }

  toggleForm(event) {
    event.preventDefault();
    this.isLogin = !this.isLogin;
    this.render();
    this.addEventListeners();
  }
}

customElements.define('gowc-auth-form', AuthFormComponent);