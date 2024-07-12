class ToastComponent extends HTMLElement {
  static get observedAttributes() {
    return ["type", "message"];
  }

  constructor() {
    super();
    this.attachShadow({ mode: "open" });
  }

  connectedCallback() {
    this.render();
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (oldValue !== newValue) {
      this.render();
    }
  }

  render() {
    const type = this.getAttribute("type") || "success";
    const message = this.getAttribute("message") || "";

    this.shadowRoot.innerHTML = `
      <div
        class="toast toast--${type}"
        id="toast"
      >
        ${message}
      </div>

      <style>
        .toast {
          padding: 1rem;
          border-radius: var(--border-radius-sm);
          color: var(--color-white);
          position: fixed;
          bottom: 1rem;
          right: 1rem;
          box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.1);
          opacity: 0;
          transition: opacity 0.3s ease-in-out;
        }
        .toast--success {
          background-color: var(--color-green-600);
        }
        .toast--error {
          background-color: var(--color-red-600);
        }
        .toast.show {
          opacity: 1;
        }
      </style>
    `;
  }

  show() {
    this.shadowRoot.querySelector(".toast").classList.add("show");
    setTimeout(() => this.hide(), 3000);
  }

  hide() {
    this.shadowRoot.querySelector(".toast").classList.remove("show");
  }
}

customElements.define("gowc-toast", ToastComponent);
