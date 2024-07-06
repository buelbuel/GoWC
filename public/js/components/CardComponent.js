class CardComponent extends HTMLElement {
  connectedCallback() {
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.innerHTML = `
        <div class="card">
          <slot></slot>
        </div>
        <style>
          .card {
            max-width: var(--breakpoint-sm);
            margin: 0 auto;
            padding: 2rem;
            border-radius: var(--border-radius-md);
            background-color: var(--color-gray-800);

            @media (prefers-color-scheme: light) {
              background-color: var(--color-gray-300);
            }
          }
        </style>
      `;
  }
}

customElements.define('gowc-card', CardComponent);