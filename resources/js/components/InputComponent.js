class InputComponent extends HTMLElement {
  connectedCallback() {
    const id = this.getAttribute("id") || "text";
    const type = this.getAttribute("type") || "text";
    const name = this.getAttribute("name") || "";
    const placeholder = this.getAttribute("placeholder") || "";

    this.attachShadow({ mode: "open" });
    this.shadowRoot.innerHTML = `
        <input
          id="${id}"
          type="${type}"
          name="${name}"
          placeholder="${placeholder}"
          required
        >

        <style>
          input {
            margin-bottom: 1rem;
            padding: 0.5rem;
            font-size: var(--font-size-sm);
            width: 100%;
            border-width: 2px;
            border-style: solid;
            border-color: var(--color-gray-500);
            border-radius: var(--border-radius-sm);
            box-sizing: border-box;
            transition: border 0.3s;
            outline: none;
            
            &:focus {
              border-color: var(--color-gray-800);
            }

            @media (prefers-color-scheme: dark) {
              border-color: var(--color-gray-700);

              &:focus {
                border-color: var(--color-gray-500);
              }
            }
          }
        </style>
      `;

    this.input = this.shadowRoot.querySelector("input");
    this.input.addEventListener("input", this.handleInput.bind(this));
  }

  handleInput() {
    this.dispatchEvent(
      new CustomEvent("gowc-input", {
        bubbles: true,
        composed: true,
        detail: {
          name: this.input.name,
          value: this.input.value,
        },
      })
    );
  }
}

customElements.define("gowc-input", InputComponent);
