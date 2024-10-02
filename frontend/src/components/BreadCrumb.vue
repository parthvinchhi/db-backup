<template>
  <ol class="right-4 absolute flex text-white">
    <li>
      <router-link to="/">Home</router-link>
    </li>
    <li>{{ currentPath }}</li>
  </ol>
</template>

<script>
export default {
  computed: {
    currentPath() {
      return this.$route.path;
    },
  },
};
</script>
