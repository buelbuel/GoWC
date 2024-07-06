class InputComponent extends HTMLElement {
    connectedCallback() {
      const type = this.getAttribute('type') || 'text';
      const name = this.getAttribute('name') || '';
      const placeholder = this.getAttribute('placeholder') || '';
      this.attachShadow({ mode: 'open' });
      this.shadowRoot.innerHTML = `
        <input type="${type}" name="${name}" placeholder="${placeholder}" required>
        <style>
          input {
            margin-bottom: 1rem;
            padding: 0.5rem;
            font-size: 1rem;
            width: 100%;
            box-sizing: border-box;
          }
        </style>
      `;
    }
  }
  
  customElements.define('gowc-input', InputComponent);