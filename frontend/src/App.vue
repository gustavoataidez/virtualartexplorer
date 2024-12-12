<template>
  <RouterView />
</template>

<script>
import { RouterView } from "vue-router";

export default {
  components: {
    RouterView,
  },
  mounted() {
    console.log("::: Virtual Art Explorer - PI 2024");
    console.log(
      "::: Developed by: @adreoluiz, @gustavoataidez e @IsaqueLemos"
    );
  },
};
</script>

<style scoped>
body,
#app {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
  width: 100%;
  max-width: 100%;
}

#app {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
  width: 100%;
  max-width: 100%;
}
</style>
