class CardHeaderComponent extends HTMLElement {
  connectedCallback() {
    this.attachShadow({ mode: "open" });
    this.shadowRoot.innerHTML = `
        <h2 class="card-header">
          <slot></slot>
        </h2>

        <style>
          .card-header {
            margin: 0 0 1rem 0;
            text-align: center;
            font-size: 1.5rem;
          }
        </style>
      `;
  }
}

customElements.define("gowc-card-header", CardHeaderComponent);
