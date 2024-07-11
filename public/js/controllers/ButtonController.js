const controllers = {
  exampleController: () => {
    console.log("Button clicked 🤓");
  },
  // Add more controllers here
};

document.addEventListener("DOMContentLoaded", () => {
  document.querySelectorAll("gowc-button").forEach((buttonComponent) => {
    const shadowRoot = buttonComponent.shadowRoot;
    const button = shadowRoot ? shadowRoot.querySelector("button") : null;

    if (button) {
      const controller = button.getAttribute("data-controller");

      button.addEventListener("click", () => {
        if (controller && controllers[controller]) {
          controllers[controller]();
        } else {
          console.log("No controllers defined");
        }
      });
    } else {
      console.log("Button not found in component:", buttonComponent);
    }
  });
});
