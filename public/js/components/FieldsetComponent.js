class FieldsetComponent extends HTMLElement {
    connectedCallback() {
      const legend = this.getAttribute('legend') || '';
      this.attachShadow({ mode: 'open' });
      this.shadowRoot.innerHTML = `
        <fieldset>
          <legend>${legend}</legend>
          <slot></slot>
        </fieldset>
        <style>
          fieldset {
            border: 1px solid #ccc;
            border-radius: 8px;
            padding: 1rem;
            margin-bottom: 1rem;
          }
          legend {
            font-weight: var(--font-weight-semibold);
            padding: 0 0.5rem;
          }
        </style>
      `;
    }
  }
  
  customElements.define('gowc-fieldset', FieldsetComponent);