
<template>
  <TheHeader />
  <router-view></router-view>
  <Footer />
</template>

<script>
import Footer from "./components/Footer.vue";
import TheHeader from "./components/TheHeader.vue";
export default {
  components: {
    TheHeader: TheHeader,
    Footer : Footer
  },
};
</script>

