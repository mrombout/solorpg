import Vue from 'vue'
import Vuex from 'vuex'
import Axios from 'axios'

Vue.use(Vuex)

const state = {
    activeStory: {
        name: 'A story name',
        events: [
            {
                type: 'say',
                text: 'Here is a piece of text.'
            },
            {
                type: 'roll',
                Result: 21,
                Dice: [
                    {
                        Faces: 6,
                        Result: 1
                    },
                    {
                        faces: 6,
                        result: 2
                    },
                    {
                        faces: 6,
                        result: 3
                    },
                    {
                        faces: 6,
                        result: 4
                    },
                    {
                        faces: 6,
                        result: 5
                    },
                    {
                        faces: 6,
                        result: 6
                    },
                ]
            },
            {
                type: 'ask',
                result: 'Yes, but...'
            },
            {
                type: 'gimme',
                result: 'A shy dog'
            }
        ]
    }
}

const mutations = {
    addSay(state, payload) {
        let events = [...state.activeStory.events]
        payload.id = new Date().getTime()
        events.push(payload)

        Vue.set(state.activeStory, 'events', events)
    },
    addRoll(state, payload) {
        let events = [...state.activeStory.events]
        payload.id = new Date().getTime()
        events.push(payload)

        Vue.set(state.activeStory, 'events', events)
    }
}

const actions = {
    addSay({commit}, payload) {
        commit('addSay', payload)
    },
    addRoll({commit}) {
        Axios
            .get('https://us-central1-cloud-gm-zargon.cloudfunctions.net/roll')
            .then(response => {
                let payload = response.data
                payload.type = 'roll'
                commit('addRoll', payload)
            })
    }
}

export default new Vuex.Store({
    state: state,
    actions: actions,
    mutations: mutations
})
