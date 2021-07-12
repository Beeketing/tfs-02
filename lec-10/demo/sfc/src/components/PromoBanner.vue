<template>
  <div class="col-xs-4 col-xxs-12 mb-30 promo-banner">
    <a href="#">
      <img
        :src="image"
        alt=""
      >
      <div class="overlay" />
      <div class="promo-inner valign">
        <h2>{{ title }}</h2>
        <span v-if="description">{{ description }}</span>
      </div>
    </a>
  </div>
</template>

<script>
export default {
  name: 'PromoBanner',
  props: {
    title: {
      type: String,
      required: true,
    },
    description: {
      type: String,
      require: false,
      default: '',
    },
    image: {
      type: String,
      required: true,
    }
  }
};
</script>

<style scoped>

</style>
