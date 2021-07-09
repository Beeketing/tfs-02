<template>
  <div class="box">
    <img :src="product.img" :alt="product.title" />
    <i class="fa fa-plus" @click="$emit('addToCart', product)" />
    <h2>{{ product.title }}</h2>
    <p>${{ product.price }}</p>
  </div>
</template>

<script>
export default {
  name: 'Product',
  props: {
    product: {
      type: Object,
      default: () => {},
    },
  },
};
</script>
