class AuthFormComponent extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
  }

  connectedCallback() {
    this.renderLoginForm();
  }

  renderLoginForm() {
    this.shadowRoot.innerHTML = `
      <gowc-card>
        <gowc-card-header title="Login"></gowc-card-header>
        <form action="/login" method="POST">
          <gowc-input type="text" name="username" placeholder="Username"></gowc-input>
          <gowc-input type="password" name="password" placeholder="Password"></gowc-input>
          <gowc-button button-type="submit" text="Login"></gowc-button>
        </form>
        <p>No account? <a href="#" id="switch-to-register">Register here</a></p>
      </gowc-card>
    `;

    this.shadowRoot.querySelector("#switch-to-register").addEventListener(
      "click",
      (event) => {
        event.preventDefault();
        this.renderRegisterForm();
      }
    );
  }

  renderRegisterForm() {
    this.shadowRoot.innerHTML = `
      <gowc-card>
        <gowc-card-header title="Register"></gowc-card-header>
        <form action="/register" method="POST">
          <gowc-input type="text" name="username" placeholder="Username"></gowc-input>
          <gowc-input type="email" name="email" placeholder="Email"></gowc-input>
          <gowc-input type="password" name="password" placeholder="Password"></gowc-input>
          <gowc-button button-type="submit" text="Register"></gowc-button>
        </form>
        <p>Already have an account? <a href="#" id="switch-to-login">Login here</a></p>
      </gowc-card>
    `;

    this.shadowRoot.querySelector("#switch-to-login").addEventListener(
      "click",
      (event) => {
        event.preventDefault();
        this.renderLoginForm();
      }
    );
  }
}

customElements.define("gowc-auth-form", AuthFormComponent);