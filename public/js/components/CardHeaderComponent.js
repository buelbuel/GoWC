class CardHeaderComponent extends HTMLElement {
  connectedCallback() {
    this.attachShadow({ mode: "open" });
    const title = this.getAttribute("title") || "";
    this.shadowRoot.innerHTML = `
        <div class="card-header">
          <h2>${title}</h2>
          <slot></slot>
        </div>
        <style>
          .card-header {
            margin-bottom: 1rem;
            text-align: center;
          }
          .card-header h2 {
            margin: 0;
            font-size: 1.5rem;
          }
        </style>
      `;
  }
}

customElements.define("gowc-card-header", CardHeaderComponent);
