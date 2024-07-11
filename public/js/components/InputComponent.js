class InputComponent extends HTMLElement {
  connectedCallback() {
    const id = this.getAttribute("id") || "text";
    const type = this.getAttribute("type") || "text";
    const name = this.getAttribute("name") || "";
    const placeholder = this.getAttribute("placeholder") || "";
    this.attachShadow({ mode: "open" });
    this.shadowRoot.innerHTML = `
        <input id="${id}" type="${type}" name="${name}" placeholder="${placeholder}" required>
        <style>
          input {
            margin-bottom: 1rem;
            padding: 0.5rem;
            font-size: var(--font-size-sm);
            width: 100%;
            border: none;
            border-radius: var(--border-radius-md);
            box-sizing: border-box;
          }
        </style>
      `;

    this.input = this.shadowRoot.querySelector('input');
    this.input.addEventListener('input', this.handleInput.bind(this));
  }

  handleInput() {
    this.dispatchEvent(new CustomEvent('gowc-input', {
      bubbles: true,
      composed: true,
      detail: {
        name: this.input.name,
        value: this.input.value
      }
    }));
  }
}

customElements.define("gowc-input", InputComponent);