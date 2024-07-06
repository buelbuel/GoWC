class NavMainComponent extends HTMLElement {
  connectedCallback() {
    this.innerHTML = `
        <nav id="nav__main" class="container">
          <div class="nav__wrapper">
            <a href="/">
              <span class="nav__logo">GoWC</span>
            </a>
            <ul class="nav__list"></ul>
          </div>
        </nav>
        
        <style>
          #nav__main {
            position: fixed;
            top: 0;
            left: 0;
            right: 0;
            z-index: 1;
            padding: 0;
            backdrop-filter: blur(1rem);
            transition: box-shadow 0.3s ease, background-color 0.3s ease;
          }
          
          @media (min-width: 656px) {
            #nav__main {
              border-radius: var(--border-radius-xl);
              top: 1rem;
            }
          }
          
          .nav__wrapper {
            padding: 0.875rem 2rem;
            display: flex;
            justify-content: space-between;
            align-items: center;
          }
          
          .nav__logo {
            width: 125px;
            font-size: 2rem;
            color: var(--color-white);
            transition: text-shadow 0.3s ease;
          }
          
          .nav__logo:hover {
            text-shadow: 0.25rem 0.25rem 0 var(--color-black);
          }
          
          .nav__list {
            display: inline-flex;
            gap: 1rem;
            justify-content: end;
            align-items: center;
            list-style: none;
          }
          
          .nav--scrolled {
            box-shadow: var(--shadow-md);
            background-color: rgba(var(--color-black-rgb), 0.7);

            @media (prefers-color-scheme: light) {
              background-color: rgba(var(--color-gray-800-rgb), 0.8);
            }
          }
        </style>
      `;

    const menuItems = JSON.parse(this.getAttribute("menu-items"));
    const navList = this.querySelector(".nav__list");
    menuItems.forEach((item) => {
      const li = document.createElement("li");
      const a = document.createElement("a");
      a.href = item.url;
      a.textContent = item.label;
      li.appendChild(a);
      navList.appendChild(li);
    });

    this.addScrollListener();
  }

  addScrollListener() {
    window.addEventListener("scroll", this.handleScroll.bind(this));
  }

  handleScroll() {
    const scrollPosition = window.scrollY;
    const nav = this.querySelector("#nav__main");
    if (scrollPosition > 50) {
      nav.classList.add("nav--scrolled");
    } else {
      nav.classList.remove("nav--scrolled");
    }
  }
}

customElements.define("gowc-nav-main", NavMainComponent);
