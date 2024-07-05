import { Controller } from 'https://unpkg.com/@hotwired/stimulus/dist/stimulus.js'

export default class extends Controller {
  static targets = ['nav']

  connect() {
    this.addScrollListener()
  }

  addScrollListener() {
    window.addEventListener('scroll', this.handleScroll.bind(this))
  }

  handleScroll() {
    const scrollPosition = window.scrollY
    if (scrollPosition > 50) {
      this.navTarget.classList.add('nav--scrolled')
    } else {
      this.navTarget.classList.remove('nav--scrolled')
    }
  }
}