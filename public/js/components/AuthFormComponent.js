class AuthFormComponent extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.isLogin = true;
    this.formData = {};
  }

  connectedCallback() {
    this.render();
    this.addEventListeners();
  }

  render() {
    const formType = this.isLogin ? "Login" : "Register";
    const switchText = this.isLogin ? "Register here" : "Login here";
    const switchId = this.isLogin ? "switch-to-register" : "switch-to-login";

    this.shadowRoot.innerHTML = `
      <gowc-card>
        <gowc-card-header>${formType}</gowc-card-header>
        <form id="auth-form">
          ${
            !this.isLogin
              ? `
              <gowc-label for="username">Username</gowc-label>
              <gowc-input id="username" name="username" type="text" placeholder="Username" required></gowc-input>
              <gowc-label for="email">Email</gowc-label>
              <gowc-input id="email" name="email" type="email" placeholder="Email" required></gowc-input>
              `
              : `
              <gowc-label for="email">Email</gowc-label>
              <gowc-input id="email" name="email" type="email" placeholder="Email" required></gowc-input>
              `
          }
          <gowc-label for="password">Password</gowc-label>
          <gowc-input id="password" name="password" type="password" placeholder="Password" required></gowc-input>
          <button type="submit">${formType}</button>
        </form>
        <p>${
          this.isLogin ? "No account?" : "Already have an account?"
        } <a href="#" id="${switchId}">${switchText}</a></p>
      </gowc-card>
    `;
  }

  addEventListeners() {
    const form = this.shadowRoot.querySelector("#auth-form");
    const switchLink = this.shadowRoot.querySelector(
      "#switch-to-register, #switch-to-login"
    );

    form.addEventListener("submit", this.handleSubmit.bind(this));
    switchLink.addEventListener("click", this.toggleForm.bind(this));
    this.shadowRoot.addEventListener("gowc-input", this.handleInput.bind(this));
  }

  handleInput(event) {
    const { name, value } = event.detail;
    this.formData[name] = value;
  }

  async handleSubmit(event) {
    event.preventDefault();

    const endpoint = this.isLogin ? "/api/login" : "/api/register";
    const toast = document.getElementById("toast");

    try {
      const response = await fetch(endpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(this.formData),
      });

      const result = await response.json();

      if (response.ok) {
        toast.setAttribute("type", "success");
        toast.setAttribute(
          "message",
          this.isLogin ? "Login successful" : "Registration successful"
        );
        toast.show();
        if (this.isLogin) {
          localStorage.setItem("token", result.token);
        }
      } else {
        toast.setAttribute("type", "error");
        toast.setAttribute(
          "message",
          result.error || "An error occurred. Please try again."
        );
        toast.show();
      }
    } catch (error) {
      toast.setAttribute("type", "error");
      toast.setAttribute("message", "An error occurred. Please try again.");
      toast.show();
    }
  }

  toggleForm(event) {
    event.preventDefault();
    this.isLogin = !this.isLogin;
    this.formData = {}; // Reset form data
    this.render();
    this.addEventListeners();
  }
}

customElements.define("gowc-auth-form", AuthFormComponent);
