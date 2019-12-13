import firebase from 'firebase'
import 'firebase/firestore'

const config = {
    apiKey: "AIzaSyDxuIsyuZgSuW3WtGj0IrXa5zIZUhERse8",
    authDomain: "cloud-gm-zargon.firebaseapp.com",
    projectId: 'cloud-gm-zargon',
}
firebase.initializeApp(config)

const db = firebase.firestore()
const auth = firebase.auth()

const usersCollection = db.collection('users')
const storiesCollection = db.collection('stories')

export {
    db,
    auth,
    usersCollection,
    storiesCollection,
}