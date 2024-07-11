class LabelComponent extends HTMLElement {
  connectedCallback() {
    const text = this.getAttribute("text") || "";
    const htmlFor = this.getAttribute("for") || "";
    this.attachShadow({ mode: "open" });
    this.shadowRoot.innerHTML = `
        <label for="${htmlFor}"><slot></slot></label>
        <style>
          label {
            display: block;
            margin-bottom: 0.5rem;
            font-size: var(--font-size-sm);
            font-weight: var(--font-weight-semibold);
          }
        </style>
      `;
  }
}

customElements.define("gowc-label", LabelComponent);