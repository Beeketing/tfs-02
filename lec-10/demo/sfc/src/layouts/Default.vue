<template>
  <main class="main-wrapper">
    <DefaultHeader />
    <div class="content-wrapper oh">
      <RouterView />
      <DefaultFooter />
    </div>
  </main>
</template>

<script>
import DefaultHeader from './DefaultHeader.vue';
import DefaultFooter from './DefaultFooter.vue';
export default {
  name: 'Default',
  components: {
    DefaultHeader,
    DefaultFooter,
  }
};
</script>

<style scoped>

</style>
