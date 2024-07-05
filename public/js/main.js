import { Application } from 'https://unpkg.com/@hotwired/stimulus/dist/stimulus.js'
import ButtonController from './controllers/button_controller.js'
import NavController from './controllers/nav_controller.js'

const application = Application.start()

const controllers = {
  button: ButtonController,
  nav: NavController,
  // Add more controllers here as needed
}

Object.entries(controllers).forEach(([name, controller]) => {
  application.register(name, controller)
})