<template>
  <component :is="layout" />
</template>
<script>
export default {
  name: 'App',
  components: {
    DefaultLayout: () => import('./layouts/Default.vue'),
    AccountLayout: () => import('./layouts/Account.vue'),
  },
  computed: {
    layout() {
      return this.$route.meta.layout ? `${this.$route.meta.layout}-layout` : 'default-layout'
    }
  }
};
</script>
