<template>
  <div
    v-if="$store.state.product.products.length > 0"
    class="product"
  >
    <product
      v-for="product in $store.state.product.products"
      :key="product.id"
      :product="product"
    />
  </div>
</template>
<script>
import Product from '../components/Product.vue';

export default {
  name: 'Home',
  components: {
    Product,
  },
  mounted() {
    this.$store.dispatch('fetchProducts');
  },
};
</script>
