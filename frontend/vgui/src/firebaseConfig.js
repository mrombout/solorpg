import firebase from 'firebase'
import 'firebase/firestore'

const config = {
    apiKey: "AIzaSyDxuIsyuZgSuW3WtGj0IrXa5zIZUhERse8",
    authDomain: "cloud-gm-zargon.firebaseapp.com",
}
firebase.initializeApp(config)

const auth = firebase.auth()

export {
    auth,
}