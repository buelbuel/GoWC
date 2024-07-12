class CardComponent extends HTMLElement {
  connectedCallback() {
    const variant = this.getAttribute("variant") || "default";
    const padding = this.getAttribute("padding") || "2rem";
    const borderRadius = this.getAttribute("border-radius") || "var(--border-radius-md)";

    this.attachShadow({ mode: "open" });
    this.shadowRoot.innerHTML = `
      <div class="card card--${variant}">
        <slot></slot>
      </div>
    
      <style>
        .card {
          padding: ${padding};
          border-radius: var(--border-radius-${borderRadius});
          background-color: var(--card-bg-color);
          border: 2px solid var(--card-border-color);
          box-shadow: 0.5rem 0.5rem 0 0 var(--card-shadow-color);
          transition: box-shadow 0.3s;
        }

        .card:hover {
          box-shadow: 0.25rem 0.25rem 0 0 var(--card-shadow-color);
        }

        :host {
          --card-bg-color: var(--color-gray-800);
          --card-border-color: var(--color-gray-700);
          --card-shadow-color: var(--color-black);
        }

        .card--default {
          --card-bg-color: var(--color-gray-800);
          --card-border-color: var(--color-gray-700);
        }

        .card--blue {
          --card-bg-color: var(--color-blue-900);
          --card-border-color: var(--color-blue-700);
        }

        .card--green {
          --card-bg-color: var(--color-green-900);
          --card-border-color: var(--color-green-700);
        }

        .card--red {
          --card-bg-color: var(--color-red-900);
          --card-border-color: var(--color-red-700);
        }

        .card--yellow {
          --card-bg-color: var(--color-yellow-900);
          --card-border-color: var(--color-yellow-700);
        }

        @media (prefers-color-scheme: light) {
          :host {
            --card-bg-color: var(--color-gray-300);
            --card-border-color: var(--color-gray-400);
          }

          .card--default {
            --card-bg-color: var(--color-gray-300);
            --card-border-color: var(--color-gray-400);
          }

          .card--blue {
            --card-bg-color: var(--color-blue-100);
            --card-border-color: var(--color-blue-300);
          }

          .card--green {
            --card-bg-color: var(--color-green-100);
            --card-border-color: var(--color-green-300);
          }

          .card--red {
            --card-bg-color: var(--color-red-100);
            --card-border-color: var(--color-red-300);
          }

          .card--yellow {
            --card-bg-color: var(--color-yellow-100);
            --card-border-color: var(--color-yellow-300);
          }
        }
      </style>
    `;
  }
}

customElements.define("gowc-card", CardComponent);
