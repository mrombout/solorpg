import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
    count: 0,
    activeStory: {
        name: 'A story name',
        events: [
            {
                type: 'say',
                text: 'Here is a piece of text.'
            },
            {
                type: 'roll',
                total: 21,
                dice: [
                    {
                        faces: 6,
                        result: 1
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
    increment(state) {
        state.count++
    },
    decrement(state) {
        state.count--
    },
    addSay(state, payload) {
        let events = [...state.activeStory.events]
        events.push(payload)

        Vue.set(state.activeStory, 'events', events)
    }
}

const actions = {
    increment: ({commit}) => commit('increment'),
    decrement: ({commit}) => commit('decrement'),
    incrementIfOdd ({commit, state}) {
        if ((state.count + 1) % 2 == 0) {
            commit('increment')
        }
    },
    incrementAsync({commit}) {
        return new Promise((resolve) => {
            setTimeout(() => {
                commit('increment')
                resolve()
            }, 1000)
        })
    },
    addSay({commit}, payload) {
        commit('addSay', payload)
    }
}

const getters = {
    evenOrOdd: state => state.count % 2 === 0 ? 'even' : 'odd'
}

export default new Vuex.Store({
    state,
    getters,
    actions,
    mutations
})
