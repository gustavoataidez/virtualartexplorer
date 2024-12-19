import { createStore } from "vuex";

const store = createStore({
  state: {
    user: JSON.parse(localStorage.getItem("user")) || null,
    token: localStorage.getItem("authToken") || null,
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
  },
  actions: {
    login({ commit }, user) {
      commit("setUser", user);
    },
    logout({ commit }) {
      commit("clearUser");
    },
  },
  getters: {
    isLoggedIn: (state) => !!state.token,
    userFirstName: (state) => (state.user ? state.user.firstname : ""),
  },
});

export default store;
