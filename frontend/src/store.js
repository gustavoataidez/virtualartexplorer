import { createStore } from "vuex";

const store = createStore({
  state: {
    user: JSON.parse(localStorage.getItem("user")) || null,
    token: localStorage.getItem("authToken") || null,
    currentMuseumId: undefined, // Estado temporário para o ID do museu
  },
  mutations: {
    setUser(state, user) {
      state.user = user;
      state.token = user.token;
      localStorage.setItem("user", JSON.stringify(user));
      localStorage.setItem("authToken", user.token);
    },
    clearUser(state) {
      state.user = null;
      state.token = null;
      localStorage.removeItem("user");
      localStorage.removeItem("authToken");
    },
    setCurrentMuseumId(state, id) {
      state.currentMuseumId = id; // Armazena o ID do museu temporariamente
    },
    clearCurrentMuseumId(state) {
      state.currentMuseumId = undefined; // Limpa o ID do museu
    },
  },
  actions: {
    login({ commit }, user) {
      commit("setUser", user);
    },
    logout({ commit }) {
      commit("clearUser");
    },
    updateMuseumId({ commit }, id) {
      commit("setCurrentMuseumId", id);
    },
    clearMuseumId({ commit }) {
      commit("clearCurrentMuseumId");
    },
  },
  getters: {
    isLoggedIn: (state) => !!state.token,
    userFirstName: (state) => (state.user ? state.user.firstname : ""),
    currentMuseumId: (state) => state.currentMuseumId, // Getter para acessar o ID do museu
  },
});

export default store;
