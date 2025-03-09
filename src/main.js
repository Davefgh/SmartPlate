import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import './assets/main.css'
import './assets/theme.css'

// Font Awesome
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { 
  faEnvelope, 
  faLock, 
  faEye, 
  faEyeSlash, 
  faUser,
  faCircleExclamation,
  faTriangleExclamation,
  faBars,
  faXmark,
  faHome,
  faCar,
  faIdCard,
  faFileContract,
  faGear,
  faSignOutAlt,
  faBell,
  faSearch,
  faChevronRight,
  faChevronLeft,
  faChevronDown,
  faChevronUp,
  faArrowUp,
  faCheckCircle,
  faExclamationCircle,
  faHistory,
  faUserCircle,
  faAddressBook,
  faEdit,
  faCheck,
  faTimes,
  faCamera,
  faArrowLeft,
  faInfoCircle,
  faHeartbeat,
  faUsers,
  faBuilding,
  faMapMarkerAlt
} from '@fortawesome/free-solid-svg-icons'
import { faGoogle } from '@fortawesome/free-brands-svg-icons'

// Add icons to the library
library.add(
  faEnvelope, 
  faLock, 
  faEye, 
  faEyeSlash, 
  faUser,
  faGoogle,
  faCircleExclamation,
  faTriangleExclamation,
  faBars,
  faUserCircle,
  faAddressBook,
  faEdit,
  faCheck,
  faTimes,
  faCamera,
  faArrowLeft,
  faInfoCircle,
  faHeartbeat,
  faUsers,
  faBuilding,
  faMapMarkerAlt,
  faXmark,
  faHome,
  faCar,
  faIdCard,
  faFileContract,
  faGear,
  faSignOutAlt,
  faBell,
  faSearch,
  faChevronRight,
  faChevronLeft,
  faChevronDown,
  faChevronUp,
  faArrowUp,
  faCheckCircle,
  faExclamationCircle,
  faHistory
)

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.component('font-awesome-icon', FontAwesomeIcon)

app.mount('#app')
