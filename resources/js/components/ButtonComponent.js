class ButtonComponent extends HTMLElement {
  connectedCallback() {
    const buttonType = this.getAttribute("button-type") || "button";
    const variant = this.getAttribute("variant") || "blue";
    const controller = this.getAttribute("controller") || "";

    this.attachShadow({ mode: "open" });
    this.shadowRoot.innerHTML = `
      <button
        type="${buttonType}"
        class="button button--${variant}"
        data-controller="${controller}"
      >
        <slot></slot>
      </button>

      <style>
        .button {
          padding: 0.5rem 1.5rem;
          font-size: 1rem;
          font-weight: var(--font-weight-semibold);
          color: var(--button-color);
          border: 2px solid var(--button-color);
          background-color: var(--color-gray-900);
          border-radius: var(--border-radius-sm);
          cursor: pointer;
          transition: background-color 0.3s, color 0.3s, box-shadow 0.3s;
          box-shadow: 0.5rem 0.5rem 0 0 var(--color-black);

          @media (prefers-color-scheme: light) {
            background-color: var(--color-white);
          }
        }

        .button:hover {
          background-color: var(--button-color);
          color: white;
          box-shadow: 0.25rem 0.25rem 0 0 var(--color-black);
        }

        :host {
          --button-color: var(--color-blue-600);
        }

        .button--blue {
          --button-color: var(--color-blue-600);
        }

        .button--green {
          --button-color: var(--color-green-600);
        }

        .button--red {
          --button-color: var(--color-red-600);
        }

        .button--yellow {
          --button-color: var(--color-yellow-600);
        }

        .button--gray {
          --button-color: var(--color-gray-600);
        }

        .button--white {
          --button-color: var(--color-white);
        }

        .button--black {
          --button-color: var(--color-black);
        }
      </style>
    `;

    const button = this.shadowRoot.querySelector("button");
    button.addEventListener("click", (event) => {
      if (buttonType === "submit") {
        event.preventDefault();
        this.dispatchEvent(
          new CustomEvent("gowc-button-click", {
            bubbles: true,
            composed: true,
          })
        );
      }
    });
  }
}

customElements.define("gowc-button", ButtonComponent);
