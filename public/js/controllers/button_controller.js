import { Controller } from 'https://unpkg.com/@hotwired/stimulus/dist/stimulus.js'

export default class extends Controller {
  connect() {
    this.element.addEventListener('click', () => {
      console.log('Button clicked 🤓')
    })
  }
}